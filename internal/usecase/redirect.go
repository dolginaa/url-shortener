package usecase

import (
	"fmt"

	"github.com/dolginaa/url-shortener/internal/domain"
)

func (us UrlShortener) Redirect(shortURL domain.ShortenedURL) (domain.OriginalURL, error) {
	originalURL, err := us.storage.GetByShort(shortURL)
	if err != nil {
		return domain.OriginalURL{}, err
	}

	if len(originalURL.OriginalURL) == 0 {
		return domain.OriginalURL{}, fmt.Errorf("original url for '%s' doesn't exist in storage", shortURL)
	}

	return originalURL, nil
}
