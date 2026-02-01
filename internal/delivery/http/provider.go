package http

import "github.com/dolginaa/url-shortener/internal/domain"

type URLShortener interface {
	Shorten(domain.OriginalURL) (domain.ShortenedURL, error)
	Redirect(domain.ShortenedURL) (domain.OriginalURL, error)
}

type Handler struct {
	URLShortener URLShortener
}

func NewHandler(urlShortener URLShortener) *Handler {
	return &Handler{URLShortener: urlShortener}
}
