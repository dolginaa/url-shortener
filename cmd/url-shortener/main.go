package main

import (
	delivery_http "github.com/dolginaa/url-shortener/internal/delivery/http"
	"github.com/dolginaa/url-shortener/pkg/http"
)

func main() {
	serverMux := delivery_http.RouteHttp()
	s := http.NewServer(serverMux)

	s.Start()
}
