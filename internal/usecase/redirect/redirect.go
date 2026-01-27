package redirect

import "github.com/dolginaa/url-shortener/internal/domain"

func Redirect(shortURL domain.ShortenedURL) (domain.OriginalURL, error) {
	return "/original-url", nil
}
