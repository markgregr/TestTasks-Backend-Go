package repository

import "time"

type Repository struct {
	BaseURL string
	Cache   *Cache
}

func NewRepository(baseURL string, ttl time.Duration) *Repository {
	return &Repository{
		BaseURL: baseURL,
		Cache:   NewCache(ttl),
	}
}