package domain

import (
	"errors"
	"fmt"
)

// OriginalNotFoundErr is an error that indicates that original URL was not found in the storage
type OriginalNotFoundErr struct {
	shortenedURL ShortenedURL
}

func NewOriginalNotFoundErr(shortenedUrl ShortenedURL) error {
	return OriginalNotFoundErr{shortenedUrl}
}

func (e OriginalNotFoundErr) Error() string {
	return fmt.Sprintf("original URL not found for %s", e.shortenedURL.ShortenedURL)
}

func IsOriginalNotFoundErr(err error) bool {
	target := OriginalNotFoundErr{}

	return errors.As(err, &target)
}

// ShortenedAlreadyExistsErr is an error that indicates that shortened URL already exists for a given original URL
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

func IsShortenedAlreadyExistsErr(err error) bool {
	target := ShortenedAlreadyExistsErr{}

	return errors.As(err, &target)
}
