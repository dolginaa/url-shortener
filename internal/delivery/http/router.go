package http

import (
	"github.com/gorilla/mux"
)

func RouteHttp() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/shorten", ShortenHttp()).Methods("GET")
	mux.Handle("/redirect", RedirectHttp()).Methods("POST")

	return mux
}
