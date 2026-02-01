package domain

import "fmt"

type ShortenedURL struct {
	ShortenedURL string `json:"shortened_url"`
}

func NewShortenedURL(sURL string) (ShortenedURL, error) {
	if len(sURL) == 0 {
		return ShortenedURL{}, fmt.Errorf("shortened url cannot be empty")
	}

	return ShortenedURL{
		ShortenedURL: sURL,
	}, nil
}

type OriginalURL struct {
	OriginalURL string `json:"original_url"`
}

func NewOriginalURL(oURL string) (OriginalURL, error) {
	if len(oURL) == 0 {
		return OriginalURL{}, fmt.Errorf("original url cannot be empty")
	}

	return OriginalURL{
		OriginalURL: oURL,
	}, nil
}

type ID int
