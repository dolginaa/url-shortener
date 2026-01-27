package http

import "net/http"

func RouteHttp() {
	mux := http.NewServeMux()

	mux.Handle("/shorten", ShortenHttp())
	mux.Handle("/redirect", RedirectHttp())

	//mux.ServeHTTP()
}
