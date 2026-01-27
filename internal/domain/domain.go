package domain

import "fmt"

type ShortenedURL string

func NewShortenedURL(sURL string) (ShortenedURL, error) {
	if len(sURL) == 0 {
		return "", fmt.Errorf("shortened url cannot be empty")
	}

	return ShortenedURL(sURL), nil
}

type OriginalURL string

func NewOriginalURL(oURL string) (OriginalURL, error) {
	if len(oURL) == 0 {
		return "", fmt.Errorf("original url cannot be empty")
	}

	return OriginalURL(oURL), nil
}

type ID int
