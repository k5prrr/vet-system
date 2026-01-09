package api

import (
	"app/internal/app/core/domain"
	"app/internal/app/core/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (r *Router) users(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, err := req.Cookie("auth_token")
	if err != nil {
		r.err(w, http.StatusUnauthorized, errors.New("missing auth token"))
		return
	}
	token := cookie.Value

	// Извлекаем ID, если есть
	path := strings.TrimPrefix(req.URL.Path, "/api/users")
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var id int64
	if len(parts) == 1 && parts[0] != "" {
		id, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid user ID"))
			return
		}
	} else if len(parts) > 1 {
		r.err(w, http.StatusBadRequest, errors.New("invalid URL"))
		return
	}

	switch req.Method {
	case http.MethodGet:
		if id == 0 {
			// GET /api/users — список
			users, err := r.useCase.Users(req.Context(), token)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, usecase.ErrUnauthorized) || errors.Is(err, usecase.ErrForbidden) {
					status = http.StatusForbidden
				}
				r.err(w, status, fmt.Errorf("list users: %w", err))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "users": users})
		} else {
			// GET /api/users/123 — один пользователь
			user, err := r.useCase.User(req.Context(), token, id)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, usecase.ErrUnauthorized) || errors.Is(err, usecase.ErrForbidden) {
					status = http.StatusForbidden
				}
				r.err(w, status, fmt.Errorf("get user: %w", err))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "user": user})
		}

	case http.MethodPost:
		if id != 0 {
			r.err(w, http.StatusBadRequest, errors.New("user ID must not be specified in POST"))
			return
		}
		var input domain.User
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
			return
		}
		newID, err := r.useCase.CreateUser(req.Context(), token, &input)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("create user: %w", err))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "id": newID})

	case http.MethodPut:
		if id == 0 {
			r.err(w, http.StatusBadRequest, errors.New("user ID required for update"))
			return
		}
		var input domain.User
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
			return
		}
		err := r.useCase.UpdateUser(req.Context(), token, id, &input)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("update user: %w", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})

	case http.MethodDelete:
		if id == 0 {
			r.err(w, http.StatusBadRequest, errors.New("user ID required for delete"))
			return
		}
		err := r.useCase.DeleteUser(req.Context(), token, id)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("delete user: %w", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})

	default:
		r.err(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
}
