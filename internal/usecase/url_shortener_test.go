package usecase

import (
	"testing"

	"github.com/dolginaa/url-shortener/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUrlShortener_Shorten(t *testing.T) {
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

	urlShortener := NewURLShortener(storageMock, aliasMap)

	tests := []struct {
		name                 string
		originalURL          domain.OriginalURL
		expectedShortenedURL domain.ShortenedURL
		assertErr            func(error) bool
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
			assertErr: func(err error) bool {
				return domain.IsShortenedAlreadyExistsErr(err)
			},
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
			t.Parallel()

			resShortenedURL, err := urlShortener.Shorten(tt.originalURL)
			if tt.assertErr != nil {
				assert.Error(t, err)
				assert.True(t, tt.assertErr(err), "unexpected error type")
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, resShortenedURL, tt.expectedShortenedURL)
		})
	}
}

func TestUrlShortener_Redirect(t *testing.T) {
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

	urlShortener := NewURLShortener(storageMock, aliasMap)

	tests := []struct {
		name                string
		shortenedURL        domain.ShortenedURL
		expectedOriginalURL domain.OriginalURL
		assertErr           func(error) bool
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
			assertErr: func(err error) bool {
				return domain.IsOriginalNotFoundErr(err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resOriginalURL, err := urlShortener.Redirect(tt.shortenedURL)
			if tt.assertErr != nil {
				assert.Error(t, err)
				assert.True(t, tt.assertErr(err), "unexpected error type")
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, resOriginalURL, tt.expectedOriginalURL)
		})
	}
}
