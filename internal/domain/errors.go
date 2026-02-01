package domain

import "fmt"

type OriginalNotFoundErr struct {
	shortenedURL ShortenedURL
}

func NewOriginalNotFoundErr(shortenedUrl ShortenedURL) error {
	return OriginalNotFoundErr{shortenedUrl}
}

func (e OriginalNotFoundErr) Error() string {
	return fmt.Sprintf("original URL not found for %s", e.shortenedURL.ShortenedURL)
}

type ShortenedAlreadyExistsErr struct {
	shortenedURL ShortenedURL
	originalURL  OriginalURL
}

func NewShortenedAlreadyExistsErr(shortenedURL ShortenedURL, originalURL OriginalURL) error {
	return ShortenedAlreadyExistsErr{shortenedURL: shortenedURL, originalURL: originalURL}
}

func (e ShortenedAlreadyExistsErr) Error() string {
	return fmt.Sprintf("shortened URL %s already exists for %s", e.shortenedURL.ShortenedURL, e.originalURL.OriginalURL)
}
