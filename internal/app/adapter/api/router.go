package api

import (
	"app/internal/app/core/port"
	"encoding/json"
	"net/http"
)

type Router struct {
	mux     *http.ServeMux
	useCase port.IUseCase
}

func New(useCase port.IUseCase) *Router {
	r := &Router{
		mux:     http.NewServeMux(),
		useCase: useCase,
	}
	r.init()
	return r
}
func (r *Router) init() {
	r.HandleFunc("/api/create_admin", r.createAdmin)
	r.HandleFunc("/api/login", r.login)
	r.HandleFunc("/api/logout", r.logout)

	//r.HandleFunc("/api/current_user", r.currentUser)
	r.HandleFunc("/api/init", r.initData)

	r.HandleFunc("/api/users", r.users)
	r.HandleFunc("/api/animals", r.animals)

	r.mux.Handle("/",
		http.StripPrefix("/",
			http.FileServer(http.Dir("./static/")),
		),
	)
}

// ---

// ---

// Прокидываем для удобства
func (r *Router) err(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":    false,
		"error": err.Error(),
	})
}
func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc(path, f)
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
