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

	// init shortenerUC
	shortenerUC := usecase.NewURLShortener(storage, nil)

	// init http handler
	httpHandler := delivery_http.NewHandler(shortenerUC)

	// init server
	s := http.NewServer(httpHandler.Route())

	// start server
	s.Start()
}
