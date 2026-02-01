package http

import (
	"github.com/gorilla/mux"
)

func (p *Handler) Route() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", p.ShortenHttp()).Methods("GET")
	router.Handle("/redirect", p.RedirectHttp()).Methods("POST")

	return router
}
