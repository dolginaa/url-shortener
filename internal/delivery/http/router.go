package http

import (
	"github.com/gorilla/mux"
)

func (p *Provider) RouteHttp() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/shorten", p.ShortenHttp()).Methods("GET")
	mux.Handle("/redirect", p.RedirectHttp()).Methods("POST")

	return mux
}
