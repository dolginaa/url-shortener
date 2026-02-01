package infrastructure

import (
	"sync"

	"github.com/dolginaa/url-shortener/internal/domain"
)

type Storage struct {
	mu   sync.Mutex
	data map[string]string
}

func NewStorage() *Storage {
	data := make(map[string]string)
	return &Storage{
		data: data,
	}
}

func (s *Storage) GetByShort(shortenedURL domain.ShortenedURL) (domain.OriginalURL, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	originalStr, found := s.data[shortenedURL.ShortenedURL]
	if !found {
		return domain.OriginalURL{}, domain.NewOriginalNotFoundErr(shortenedURL)
	}
	return domain.OriginalURL{OriginalURL: originalStr}, nil
}

func (s *Storage) Save(originalURL domain.OriginalURL, shortenedURL domain.ShortenedURL) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	originalStrOld, found := s.data[shortenedURL.ShortenedURL]
	if found && originalStrOld != originalURL.OriginalURL {
		return domain.NewShortenedAlreadyExistsErr(shortenedURL, originalURL)
	}

	for short, o := range s.data {
		if o == originalURL.OriginalURL && short != shortenedURL.ShortenedURL {
			delete(s.data, short)
			s.data[shortenedURL.ShortenedURL] = originalURL.OriginalURL
			return nil
		}
	}

	s.data[shortenedURL.ShortenedURL] = originalURL.OriginalURL

	return nil
}
