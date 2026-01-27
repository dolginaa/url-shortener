package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dolginaa/url-shortener/internal/domain"
	"github.com/dolginaa/url-shortener/internal/usecase/redirect"
	"github.com/dolginaa/url-shortener/internal/usecase/shorten"
)

func ShortenHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); contentType != "json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		originalReq, err := r.GetBody()
		if err != nil {
			internalError(w)
			return
		}
		defer func() {
			err := originalReq.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		var originalURL domain.OriginalURL
		if err = json.NewDecoder(originalReq).Decode(&originalURL); err != nil {
			internalError(w)
			return
		}

		shortenedURL, err := shorten.Shorten(originalURL)
		if err != nil {
			internalError(w)
			return
		}

		if err = json.NewEncoder(w).Encode(shortenedURL); err != nil {
			internalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func RedirectHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); contentType != "json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		originalReq, err := r.GetBody()
		if err != nil {
			internalError(w)
			return
		}
		defer func() {
			err := originalReq.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		var shortURL domain.ShortenedURL
		json.NewDecoder(originalReq).Buffered()
		if err = json.NewDecoder(originalReq).Decode(&shortURL); err != nil {
			internalError(w)
			return
		}

		originalURL, err := redirect.Redirect(shortURL)
		if err != nil {
			internalError(w)
			return
		}

		if err = json.NewEncoder(w).Encode(originalURL); err != nil {
			internalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
