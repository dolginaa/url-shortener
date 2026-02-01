package main

import (
	delivery_http "github.com/dolginaa/url-shortener/internal/delivery/http"
	"github.com/dolginaa/url-shortener/internal/infrastructure"
	"github.com/dolginaa/url-shortener/internal/usecase"
	"github.com/dolginaa/url-shortener/pkg/http"
)

func main() {
	// init storage
	storage := infrastructure.NewStorage()

	// init usecase
	usecase := usecase.NewUrlShortener(storage, nil)

	// init http provider
	httpProvider := delivery_http.NewProvider(usecase)

	// init server
	s := http.NewServer(httpProvider.RouteHttp())

	// start server
	s.Start()
}
