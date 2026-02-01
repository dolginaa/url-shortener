package http

import "github.com/dolginaa/url-shortener/internal/domain"

type Usecase interface {
	Shorten(domain.OriginalURL) (domain.ShortenedURL, error)
	Redirect(domain.ShortenedURL) (domain.OriginalURL, error)
}

type Provider struct {
	Usecase Usecase
}

func NewProvider(usecase Usecase) *Provider {
	return &Provider{Usecase: usecase}
}
