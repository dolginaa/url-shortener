package infrastructure

import (
	"github.com/dolginaa/url-shortener/internal/domain"
)

type Storage struct {
	data map[string]string
}

func NewStorage() *Storage {
	data := make(map[string]string)
	return &Storage{
		data: data,
	}
}

func (s *Storage) GetByShort(shortenedURL domain.ShortenedURL) (domain.OriginalURL, error) {
	originalStr, found := s.data[shortenedURL.ShortenedURL]
	if !found {
		return domain.OriginalURL{}, domain.NewOriginalNotFoundErr(shortenedURL)
	}
	return domain.OriginalURL{OriginalURL: originalStr}, nil
}

func (s *Storage) Save(originalURL domain.OriginalURL, shortenedURL domain.ShortenedURL) error {
	originalStrOld, found := s.data[shortenedURL.ShortenedURL]
	if found && originalStrOld != originalURL.OriginalURL {
		return domain.NewShortenedAlreadyExistsErr(shortenedURL, originalURL)
	}

	for s, o := range s.data {
		if o == originalURL.OriginalURL && s != shortenedURL.ShortenedURL {
			s = shortenedURL.ShortenedURL
			return nil
		}
	}

	s.data[shortenedURL.ShortenedURL] = originalURL.OriginalURL

	return nil
}
