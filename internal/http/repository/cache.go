package repository

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache представляет кеш для хранения данных
type Cache struct {
    cache *cache.Cache // экземпляр кеша из библиотеки go-cache
}

// NewCache создает новый объект кеша с заданным временем жизни (TTL)
func NewCache(ttl time.Duration) *Cache {
    return &Cache{
        cache: cache.New(ttl, ttl),
    }
}

// Get получает значение из кеша по ключу
// возвращает значение и флаг, указывающий, было ли значение найдено в кеше
func (c *Cache) Get(key string) (interface{}, bool) {
    return c.cache.Get(key) 
}

// Set устанавливает значение в кеше для указанного ключа с заданным expiration
func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
    c.cache.Set(key, value, expiration)
}
