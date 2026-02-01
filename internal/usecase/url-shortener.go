package usecase

import "github.com/dolginaa/url-shortener/internal/domain"

type UrlShortener struct {
	storage ShortenerStorage

	aliasMap map[string]string
}

var defaultAliasMap = map[string]string{
	"long-string":          "short",
	"custom-params-string": "test",
}

func NewUrlShortener(storage ShortenerStorage, aliasMap map[string]string) UrlShortener {
	if len(aliasMap) == 0 {
		aliasMap = defaultAliasMap
	}

	return UrlShortener{
		storage:  storage,
		aliasMap: aliasMap,
	}
}

type ShortenerStorage interface {
	GetByShort(domain.ShortenedURL) (domain.OriginalURL, error)
	Save(domain.OriginalURL, domain.ShortenedURL) error
}
