package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dolginaa/url-shortener/internal/domain"
)

func (p *Provider) ShortenHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		originalReq := r.Body
		defer func() {
			err := originalReq.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		var originalURL domain.OriginalURL
		if err := json.NewDecoder(originalReq).Decode(&originalURL); err != nil {
			internalError(w)
			return
		}

		shortenedURL, err := p.Usecase.Shorten(originalURL)
		if err != nil {
			if domain.IsShortenedAlreadyExistsErr(err) {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			internalError(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(shortenedURL); err != nil {
			internalError(w)
			return
		}
	}
}

func (p *Provider) RedirectHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		originalReq := r.Body
		defer func() {
			err := originalReq.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		var shortURL domain.ShortenedURL
		if err := json.NewDecoder(originalReq).Decode(&shortURL); err != nil {
			internalError(w)
			return
		}

		originalURL, err := p.Usecase.Redirect(shortURL)
		if err != nil {
			if domain.IsOriginalNotFoundErr(err) {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			internalError(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(originalURL); err != nil {
			internalError(w)
			return
		}
	}
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
