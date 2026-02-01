package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dolginaa/url-shortener/internal/domain"
)

type shortenRequest struct {
	OriginalURL string `json:"original_url"`
}

func (p *Handler) ShortenHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body := r.Body
		defer func() {
			if err := body.Close(); err != nil {
				log.Print(err)
			}
		}()

		var req shortenRequest
		if err := json.NewDecoder(body).Decode(&req); err != nil {
			internalError(w)
			return
		}

		originalURL, err := domain.NewOriginalURL(req.OriginalURL)
		if err != nil {
			badRequest(w)
			return
		}

		shortenedURL, err := p.URLShortener.Shorten(originalURL)
		if err != nil {
			if domain.IsShortenedAlreadyExistsErr(err) {
				badRequest(w)
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

type redirectRequest struct {
	ShortenedURL string `json:"shortened_url"`
}

func (p *Handler) RedirectHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body := r.Body
		defer func() {
			if err := body.Close(); err != nil {
				log.Print(err)
			}
		}()

		var req redirectRequest
		if err := json.NewDecoder(body).Decode(&req); err != nil {
			internalError(w)
			return
		}

		shortURL, err := domain.NewShortenedURL(req.ShortenedURL)
		if err != nil {
			badRequest(w)
			return
		}

		originalURL, err := p.URLShortener.Redirect(shortURL)
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

func badRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"invalid request"}`))
}

func internalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}
