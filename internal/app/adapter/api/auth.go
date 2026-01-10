package api

import (
	"app/pkg/utilities"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (r *Router) createAdmin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.useCase.CreateAdmin(req.Context())
	if err != nil {
		r.err(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok": true,
	})
}
func (r *Router) login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodPost {
		r.err(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	var input struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
		return
	}

	input.Phone = utilities.OnlyDigits(input.Phone)
	if len(input.Phone) < 1 {
		r.err(w, http.StatusBadRequest, errors.New("invalid phone"))
		return
	}
	if input.Password == "" {
		r.err(w, http.StatusBadRequest, errors.New("invalid password"))
		return
	}

	token, err := r.useCase.Login(req.Context(), input.Phone, input.Password, req.RemoteAddr)
	if err != nil {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("login failed: %w", err))
		return
	}

	data, err := r.useCase.InitData(req.Context(), token)
	if err != nil {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: %w", err))
		return
	}

	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   int(30 * 24 * 60 * 60), // 30 дней в секундах
		HttpOnly: true,                   // JS не сможет прочитать!
		//Secure:   true,      // true в проде (только по HTTPS)
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":   true,
		"data": data,
	})
}
func (r *Router) logout(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		r.err(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	// Удаляем куку auth_token
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "", // пустое значение
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure:  true, // только по HTTPS
	})

	// Ответ: 204 No Content — стандарт для logout без тела
	//w.WriteHeader(http.StatusNoContent)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok": true,
	})
}
func (r *Router) currentUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, err := req.Cookie("auth_token")
	if err != nil {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: missing auth token: %w", err))
		return
	}

	token := cookie.Value
	if token == "" {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: missing auth token: %w", err))
		return
	}

	user, err := r.useCase.CurrentUser(req.Context(), token)
	if err != nil {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":   true,
		"user": user,
	})
}
func (r *Router) initData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, err := req.Cookie("auth_token")
	if err != nil {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: missing auth token: %w", err))
		return
	}

	token := cookie.Value
	if token == "" {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: missing auth token: %w", err))
		return
	}

	data, err := r.useCase.InitData(req.Context(), token)
	if err != nil {
		r.err(w, http.StatusUnauthorized, fmt.Errorf("unauthorized: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":   true,
		"data": data,
	})
}
