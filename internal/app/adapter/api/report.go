// ./internal/app/adapter/api/report.go
package api

import (
	"app/internal/app/core/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (r *Router) report(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method != http.MethodGet {
		r.err(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	cookie, err := req.Cookie("auth_token")
	if err != nil {
		r.err(w, http.StatusUnauthorized, errors.New("missing auth token"))
		return
	}

	report, err := r.useCase.Report(req.Context(), cookie.Value)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrUnauthorized) || errors.Is(err, usecase.ErrForbidden) {
			status = http.StatusForbidden
		}
		r.err(w, status, fmt.Errorf("get report: %w", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "report": report})
}
