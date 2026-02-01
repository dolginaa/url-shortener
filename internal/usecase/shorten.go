package usecase

import (
	"strings"

	"github.com/dolginaa/url-shortener/internal/domain"
)

func (us URLShortener) Shorten(originalURL domain.OriginalURL) (domain.ShortenedURL, error) {
	shortenedURL := us.calcShortenedURL(originalURL)

	foundByShort, err := us.storage.GetByShort(shortenedURL)
	if err != nil && !domain.IsOriginalNotFoundErr(err) {
		return domain.ShortenedURL{}, err
	}

	if len(foundByShort.OriginalURL) != 0 && foundByShort.OriginalURL != originalURL.OriginalURL {
		return domain.ShortenedURL{}, domain.NewShortenedAlreadyExistsErr(shortenedURL, originalURL)
	}

	if err = us.storage.Save(originalURL, shortenedURL); err != nil {
		return domain.ShortenedURL{}, err
	}

	return shortenedURL, nil
}

func (us URLShortener) calcShortenedURL(originalURL domain.OriginalURL) domain.ShortenedURL {
	resultURL := originalURL.OriginalURL
	for originalSubString, alias := range us.aliasMap {
		resultURL = strings.ReplaceAll(resultURL, originalSubString, alias)
	}

	return domain.ShortenedURL{ShortenedURL: resultURL}
}
