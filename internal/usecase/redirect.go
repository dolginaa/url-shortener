package usecase

import (
	"github.com/dolginaa/url-shortener/internal/domain"
)

func (us URLShortener) Redirect(shortURL domain.ShortenedURL) (domain.OriginalURL, error) {
	originalURL, err := us.storage.GetByShort(shortURL)
	if err != nil {
		return domain.OriginalURL{}, err
	}

	if len(originalURL.OriginalURL) == 0 {
		return domain.OriginalURL{}, domain.NewOriginalNotFoundErr(shortURL)
	}

	return originalURL, nil
}
