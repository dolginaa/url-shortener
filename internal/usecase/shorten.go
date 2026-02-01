package usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dolginaa/url-shortener/internal/domain"
)

func (us UrlShortener) Shorten(originalURL domain.OriginalURL) (domain.ShortenedURL, error) {
	shortenedUrl := us.calcShortenedUrl(originalURL)

	foundByShort, err := us.storage.GetByShort(shortenedUrl)
	if err != nil && errors.Is(err, domain.NewOriginalNotFoundErr(shortenedUrl)) {
		return domain.ShortenedURL{}, err
	}

	if len(foundByShort.OriginalURL) != 0 {
		return domain.ShortenedURL{}, fmt.Errorf("couldn't calculate shortened url for %s, already exists", originalURL)
	}

	if err = us.storage.Save(originalURL, shortenedUrl); err != nil {
		return domain.ShortenedURL{}, err
	}

	return shortenedUrl, nil
}

func (us UrlShortener) calcShortenedUrl(originalURL domain.OriginalURL) domain.ShortenedURL {
	var resultUrl string
	for originalSubString, alias := range us.aliasMap {
		resultUrl = strings.ReplaceAll(originalURL.OriginalURL, originalSubString, alias)
	}

	return domain.ShortenedURL{ShortenedURL: resultUrl}
}
