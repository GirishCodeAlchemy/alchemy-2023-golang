package solutions

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// URLStore holds shortened URLs and their original mappings
type URLStore struct {
	sync.RWMutex
	urlMap map[string]string
}

// NewURLStore initializes a new URL store
func NewURLStore() *URLStore {
	return &URLStore{
		urlMap: make(map[string]string),
	}
}

// ShortenURL generates a random short path for the given URL
func (s *URLStore) ShortenURL(longURL string) string {
	s.Lock()
	defer s.Unlock()

	shortURL := generateShortURL()
	s.urlMap[shortURL] = longURL
	return shortURL
}

// GetOriginalURL retrieves the original URL for the given short path
func (s *URLStore) GetOriginalURL(shortURL string) (string, bool) {
	s.RLock()
	defer s.RUnlock()

	longURL, exists := s.urlMap[shortURL]
	return longURL, exists
}

// generateShortURL generates a random 6-character string for the shortened URL
func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(shortURL)
}

// Handler for shortening URLs
func shortenHandler(store *URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}

		shortURL := store.ShortenURL(req.URL)
		response := fmt.Sprintf("http://localhost:8080/%s", shortURL)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"shortened_url": response})
	}
}

// Handler for redirecting shortened URLs to their original counterparts
func redirectHandler(store *URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := strings.TrimPrefix(r.URL.Path, "/")
		if longURL, exists := store.GetOriginalURL(shortURL); exists {
			http.Redirect(w, r, longURL, http.StatusFound)
			return
		}
		http.NotFound(w, r)
	}
}

func init() {
	store := NewURLStore()

	http.HandleFunc("/shorten", shortenHandler(store))
	http.HandleFunc("/", redirectHandler(store))

	fmt.Println("URL Shortener running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
