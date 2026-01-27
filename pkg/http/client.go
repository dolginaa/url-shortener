package http

import (
	"net/http"
)

func CreateServer() {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {})
}
