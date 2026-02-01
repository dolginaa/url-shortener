package usecase

import (
	"testing"

	"github.com/dolginaa/url-shortener/internal/domain"
	"github.com/go-jose/go-jose/v4/testutils/assert"
)

func TestShortenUrl(t *testing.T) {
	t.Parallel()

	getByShortMockRes := map[string]StorageGetMockResult{
		"github.com/adolgina/short-string": {
			originalURL: domain.OriginalURL{},
			err: domain.NewOriginalNotFoundErr(domain.ShortenedURL{
				ShortenedURL: "github.com/adolgina/short-string",
			}),
		},
		"github.com/adolgina/after": {
			originalURL: domain.OriginalURL{OriginalURL: "different-url"},
			err: domain.NewShortenedAlreadyExistsErr(
				domain.ShortenedURL{
					ShortenedURL: "github.com/adolgina/after",
				},
				domain.OriginalURL{
					OriginalURL: "github.com/adolgina/before",
				},
			),
		},
		"github.com/adolgina/compact": {
			originalURL: domain.OriginalURL{OriginalURL: "github.com/adolgina/long:with_params"},
			err:         nil,
		},
	}

	saveMockRes := map[string]error{
		"github.com/adolgina/long-string":      nil,
		"github.com/adolgina/long:with_params": nil,
	}

	aliasMap := map[string]string{
		"long-string":      "short-string",
		"before":           "after",
		"long:with_params": "compact",
	}

	storageMock := NewStorageMock(getByShortMockRes, saveMockRes)

	urlShortener := NewUrlShortener(storageMock, aliasMap)

	tests := []struct {
		name                 string
		originalURL          domain.OriginalURL
		expectedShortenedURL domain.ShortenedURL
		expectedErr          error
	}{
		{
			name: "valid shortening - ok",
			originalURL: domain.OriginalURL{
				OriginalURL: "github.com/adolgina/long-string",
			},
			expectedShortenedURL: domain.ShortenedURL{
				ShortenedURL: "github.com/adolgina/short-string",
			},
		},
		{
			name: "shortened already exists - err",
			originalURL: domain.OriginalURL{
				OriginalURL: "github.com/adolgina/before",
			},
			expectedErr: domain.NewShortenedAlreadyExistsErr(
				domain.ShortenedURL{
					ShortenedURL: "github.com/adolgina/after",
				},
				domain.OriginalURL{
					OriginalURL: "github.com/adolgina/before",
				},
			),
		},
		{
			name: "short exists for same original - ok",
			originalURL: domain.OriginalURL{
				OriginalURL: "github.com/adolgina/long:with_params",
			},
			expectedShortenedURL: domain.ShortenedURL{
				ShortenedURL: "github.com/adolgina/compact",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resShourenedURL, err := urlShortener.Shorten(tt.originalURL)
			if tt.expectedErr != nil {
				assert.Error(t, err, tt.expectedErr.Error())
				return
			}
			assert.Equal(t, resShourenedURL, tt.expectedShortenedURL)
		})
	}
}

func TestRedirectUrl(t *testing.T) {
	t.Parallel()

	getByShortMockRes := map[string]StorageGetMockResult{
		"github.com/adolgina/short-string": {
			originalURL: domain.OriginalURL{OriginalURL: "github.com/adolgina/long-string"},
			err:         nil,
		},
		"github.com/adolgina/after": {
			originalURL: domain.OriginalURL{},
			err: domain.NewOriginalNotFoundErr(
				domain.ShortenedURL{
					ShortenedURL: "github.com/adolgina/after",
				},
			),
		},
	}

	aliasMap := map[string]string{
		"long-string":      "short-string",
		"before":           "after",
		"long:with_params": "compact",
	}

	storageMock := NewStorageMock(getByShortMockRes, nil)

	urlShortener := NewUrlShortener(storageMock, aliasMap)

	tests := []struct {
		name                string
		shortenedURL        domain.ShortenedURL
		expectedOriginalURL domain.OriginalURL
		expectedErr         error
	}{
		{
			name: "valid redirect - ok",
			shortenedURL: domain.ShortenedURL{
				ShortenedURL: "github.com/adolgina/short-string",
			},
			expectedOriginalURL: domain.OriginalURL{
				OriginalURL: "github.com/adolgina/long-string",
			},
		},
		{
			name: "original not found - err",
			shortenedURL: domain.ShortenedURL{
				ShortenedURL: "github.com/adolgina/after",
			},
			expectedErr: domain.NewOriginalNotFoundErr(
				domain.ShortenedURL{
					ShortenedURL: "github.com/adolgina/after",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resOriginalURL, err := urlShortener.Redirect(tt.shortenedURL)
			if tt.expectedErr != nil {
				assert.Error(t, err, tt.expectedErr.Error())
				return
			}
			assert.Equal(t, resOriginalURL, tt.expectedOriginalURL)
		})
	}
}
