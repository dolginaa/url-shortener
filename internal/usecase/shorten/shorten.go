package shorten

import "github.com/dolginaa/url-shortener/internal/domain"

func Shorten(originalURL domain.OriginalURL) (domain.ShortenedURL, error) {
	return "short-url", nil
}
