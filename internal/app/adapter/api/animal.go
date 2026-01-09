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

func (r *Router) animals(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем токен
	cookie, err := req.Cookie("auth_token")
	if err != nil {
		r.err(w, http.StatusUnauthorized, errors.New("missing auth token"))
		return
	}
	token := cookie.Value

	path := strings.TrimPrefix(req.URL.Path, "/api/animals")
	parts := strings.Split(strings.Trim(path, "/"), "/")

	// Если есть ID в пути: /api/animals/123
	var id int64
	if len(parts) == 1 && parts[0] != "" {
		id, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid animal ID"))
			return
		}
	} else if len(parts) > 1 {
		r.err(w, http.StatusBadRequest, errors.New("invalid URL"))
		return
	}

	switch req.Method {
	case http.MethodGet:
		if id == 0 {
			// GET /api/animals — список
			animals, err := r.useCase.Animals(req.Context(), token)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, usecase.ErrUnauthorized) {
					status = http.StatusUnauthorized
				} else if errors.Is(err, usecase.ErrForbidden) {
					status = http.StatusForbidden
				}
				r.err(w, status, fmt.Errorf("get animals: %w", err))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "animals": animals})
		} else {
			// GET /api/animals/123 — один
			animal, err := r.useCase.Animal(req.Context(), token, id)
			if err != nil {
				status := http.StatusInternalServerError
				if errors.Is(err, usecase.ErrUnauthorized) || errors.Is(err, usecase.ErrForbidden) {
					status = http.StatusForbidden
				}
				r.err(w, status, fmt.Errorf("get animal: %w", err))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "animal": animal})
		}

	case http.MethodPost:
		// POST /api/animals — создание
		var input domain.Animal
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
			return
		}
		id, err := r.useCase.CreateAnimal(req.Context(), token, &input)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("create animal: %w", err))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "id": id})

	case http.MethodPut:
		if id == 0 {
			r.err(w, http.StatusBadRequest, errors.New("animal ID required for update"))
			return
		}
		// PUT /api/animals/123 — обновление
		var input domain.Animal
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			r.err(w, http.StatusBadRequest, errors.New("invalid JSON"))
			return
		}
		err := r.useCase.UpdateAnimal(req.Context(), token, id, &input)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("update animal: %w", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})

	case http.MethodDelete:
		if id == 0 {
			r.err(w, http.StatusBadRequest, errors.New("animal ID required for delete"))
			return
		}
		// DELETE /api/animals/123
		err := r.useCase.DeleteAnimal(req.Context(), token, id)
		if err != nil {
			status := http.StatusInternalServerError
			if errors.Is(err, usecase.ErrUnauthorized) {
				status = http.StatusUnauthorized
			} else if errors.Is(err, usecase.ErrForbidden) {
				status = http.StatusForbidden
			}
			r.err(w, status, fmt.Errorf("delete animal: %w", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})

	default:
		r.err(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
}
