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
	// Change
	nameEntity, nameEntities := "user", "users"
	w.Header().Set("Content-Type", "application/json")

	cookie, err := req.Cookie("auth_token")
	if err != nil {
		r.err(w, http.StatusUnauthorized, errors.New("missing auth token"))
		return
	}
	token := cookie.Value

	path := strings.TrimPrefix(req.URL.Path, fmt.Sprintf("/api/%s", nameEntities))
	parts := strings.Split(strings.Trim(path, "/"), "/")

	var id int64
	if len(parts) == 1 && parts[0] != "" {
		id, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid ID"))
			return
		}
	} else if len(parts) > 1 {
		r.err(w, http.StatusBadRequest, errors.New("invalid URL"))
		return
	}

	switch req.Method {
	case http.MethodGet:
		if id == 0 {
			// Change
			entities, err := r.useCase.Users(req.Context(), token)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, usecase.ErrUnauthorized) {
					status = http.StatusUnauthorized
				} else if errors.Is(err, usecase.ErrForbidden) {
					status = http.StatusForbidden
				}
				r.err(w, status, fmt.Errorf("get: %w", err))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, nameEntities: entities})
		} else {
			// Change
			entity, err := r.useCase.User(req.Context(), token, id)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, usecase.ErrUnauthorized) || errors.Is(err, usecase.ErrForbidden) {
					status = http.StatusForbidden
				}
				r.err(w, status, fmt.Errorf("get: %w", err))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, nameEntity: entity})
		}

	case http.MethodPost:
		// Change
		var input domain.User
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
			return
		}
		// Change
		id, err := r.useCase.CreateUser(req.Context(), token, &input)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("create: %w", err))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "id": id})

	case http.MethodPut:
		if id == 0 {
			r.err(w, http.StatusBadRequest, errors.New("ID required for update"))
			return
		}
		// Change
		var input domain.User
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
			return
		}
		// Change
		err := r.useCase.UpdateUser(req.Context(), token, id, &input)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("update: %w", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})

	case http.MethodDelete:
		if id == 0 {
			r.err(w, http.StatusBadRequest, errors.New("ID required for delete"))
			return
		}
		// Change
		err := r.useCase.DeleteUser(req.Context(), token, id)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("delete: %w", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})

	default:
		r.err(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
}
