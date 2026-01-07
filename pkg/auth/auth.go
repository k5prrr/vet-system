package auth

import (
	"bytes"
	"crypto/hmac"
	crand "crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type IAuth interface {
	Token(id int64, secret, ip string) string
	UserID(token, secret string) (int64, error)
	Secret(nBytes int) (string, error)
	Code() (int64, error)
	PasswordHash(password string) (string, error)
	PasswordVerify(password, encodedHash string) bool
}
type Auth struct {
	signing []byte
	ttl     time.Duration
}

func NewAuth(signing string, ttl time.Duration) *Auth {
	if ttl == 0 {
		ttl = 30 * 24 * time.Hour
	}

	return &Auth{
		signing: []byte(signing),
		ttl:     ttl,
	}
}

// Генерация секрета
func (a *Auth) Secret(nBytes int) (string, error) {
	b := make([]byte, nBytes)
	if _, err := crand.Read(b); err != nil {
		return "", fmt.Errorf("Secret err: %w", err)
	}

	return hex.EncodeToString(b), nil
}
func (a *Auth) hash(secret, data []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(data)

	return h.Sum(nil)
}

// Создание двойного токена
func (a *Auth) Token(id int64, secret, ip string) string {
	payload := fmt.Sprintf(
		"%d|%d|%s",
		id,
		time.Now().Add(a.ttl).Unix(),
		ip,
	)
	payloadByte := []byte(payload)

	tokenRaw := fmt.Sprintf(
		"%s|%s|%s",
		payload,
		base64.RawURLEncoding.EncodeToString(a.hash(a.signing, payloadByte)),
		base64.RawURLEncoding.EncodeToString(a.hash([]byte(secret), payloadByte)),
	)

	return base64.RawURLEncoding.EncodeToString([]byte(tokenRaw))
}

// Генерация простого кода цифр
func (a *Auth) Code() (int64, error) {
	mMax := big.NewInt(900000)
	n, err := crand.Int(crand.Reader, mMax)
	if err != nil {
		return 0, err
	}
	return n.Int64() + 100000, nil
}

// Получение ид пользователя если токен верен
// Если секрета нет, то по подписи
func (a *Auth) UserID(token, secret string) (int64, error) {
	var err error

	data, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return 0, fmt.Errorf("invalid token encoding: %w", err)
	}

	parts := bytes.Split(data, []byte{'|'})
	if len(parts) != 5 {
		return 0, fmt.Errorf("invalid token format")
	}

	IDRaw, expRaw, ip, sigB64, secretB64 := parts[0], parts[1], parts[2], parts[3], parts[4]

	exp, err := strconv.ParseInt(string(expRaw), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid expiration time: %w", err)
	}

	if time.Now().Unix() > exp {
		return 0, fmt.Errorf("token expired")
	}

	id, err := strconv.ParseInt(string(IDRaw), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid auth id: %w", err)
	}

	payload := fmt.Sprintf("%d|%d|%s", id, exp, ip)
	payloadByte := []byte(payload)

	newSig := a.hash(a.signing, payloadByte)
	oldSig, err := base64.RawURLEncoding.DecodeString(string(sigB64))
	if err != nil {
		return 0, fmt.Errorf("invalid signature encoding: %w", err)
	}
	if !hmac.Equal(newSig, oldSig) {
		return 0, fmt.Errorf("invalid token signature")
	}

	if secret != "" {
		newSig = a.hash([]byte(secret), payloadByte)
		oldSig, err = base64.RawURLEncoding.DecodeString(string(secretB64))
		if err != nil {
			return 0, fmt.Errorf("invalid signature encoding: %w", err)
		}
		if !hmac.Equal(newSig, oldSig) {
			return 0, fmt.Errorf("invalid token signature")
		}
	}

	return id, nil
}

// PBKDF2-HMAC-SHA256 (стандартный RFC 2898, 32-байтный хеш, 16-байтная соль, 100_000 итераций)
func (a *Auth) PasswordHash(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := crand.Read(salt); err != nil {
		return "", err
	}

	// Параметры: 100 000 итераций — безопасный минимум на 2026 год
	// 32 байта = 256 бит — размер SHA256
	hash := pbkdf2HMACSHA256([]byte(password), salt, 100_000, 32)

	// Сохраняем: соль + хеш (в base64)
	out := make([]byte, 0, 16+32)
	out = append(out, salt...)
	out = append(out, hash...)
	return base64.RawStdEncoding.EncodeToString(out), nil
}

func (a *Auth) PasswordVerify(password, encodedHash string) bool {
	decoded, err := base64.RawStdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false
	}
	if len(decoded) != 16+32 {
		return false
	}

	salt, storedHash := decoded[:16], decoded[16:]

	computedHash := pbkdf2HMACSHA256([]byte(password), salt, 100_000, 32)

	return subtle.ConstantTimeCompare(computedHash, storedHash) == 1
}

// pbkdf2HMACSHA256 реализует PBKDF2 с HMAC-SHA256
// lenOut — длина выходного ключа в байтах (рекомендуется 32)
func pbkdf2HMACSHA256(password, salt []byte, iter, lenOut int) []byte {
	if iter <= 0 {
		iter = 1
	}
	if lenOut <= 0 {
		lenOut = 32
	}

	const hLen = sha256.Size        // 32
	c := (lenOut + hLen - 1) / hLen // ceil(lenOut / hLen)

	// Предвыделяем итоговый буфер
	out := make([]byte, lenOut)

	// Буферы для промежуточных значений — выделяем один раз
	// ui и t совместно используют один буфер размером hLen*2
	buf := make([]byte, hLen*2)
	ui := buf[:hLen] // текущее U_j
	t := buf[hLen:]  // T_i (накопитель XOR)

	// Один экземпляр HMAC — переиспользуем
	h := hmac.New(sha256.New, password)

	// Буфер для salt || i (i — 4 байта BE)
	// Выделяем с запасом один раз (макс. длина salt + 4)
	saltIBuf := make([]byte, len(salt)+4)
	copy(saltIBuf, salt)

	for block := 1; block <= c; block++ {
		// Формируем salt || uint32_be(block)
		i := block
		saltIBuf[len(salt)] = byte(i >> 24)
		saltIBuf[len(salt)+1] = byte(i >> 16)
		saltIBuf[len(salt)+2] = byte(i >> 8)
		saltIBuf[len(salt)+3] = byte(i)

		// U1 = HMAC(password, salt || i)
		h.Reset()
		h.Write(saltIBuf)
		h.Sum(ui[:0]) // записываем прямо в ui (без аллокации)

		// T = U1
		copy(t, ui)

		// Uj = HMAC(password, Uj-1), j = 2..iter
		for j := 2; j <= iter; j++ {
			h.Reset()
			h.Write(ui)
			h.Sum(ui[:0]) // перезаписываем ui на месте

			// T ^= Uj
			for k := range t {
				t[k] ^= ui[k]
			}
		}

		// Копируем нужную часть t в out
		offset := (block - 1) * hLen
		copyLen := hLen
		if offset+copyLen > lenOut {
			copyLen = lenOut - offset
		}
		copy(out[offset:], t[:copyLen])
	}

	return out
}
