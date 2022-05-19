package cache

import (
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

// Fetcher defines methods used to obtain data from provider.
type Fetcher interface {
	// Fetch obtains string representation of id from provider.
	// It should be used as fallback if cache does not already contain the id.
	Fetch(id int) (string, error)
	// FetchAll data from provider.
	FetchAll() (map[int]string, error)
}

// Cacher defines methods used to obtain data from cache.
type Cacher interface {

	// Get returns id translated to string; boolean for cache hit/miss.
	Get(id int) (string, bool)
}

type Cache struct {
	fs Fetcher
	c  *cache.Cache
}

func NewCache(fs Fetcher) *Cache {
	return &Cache{
		fs: fs,
		c:  cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (r *Cache) Get(id int) (string, bool) {
	convertedVal := strconv.Itoa(id)

	cachedValue, found := r.c.Get(convertedVal)
	if found {
		return cachedValue.(string), true
	} else {
		value, err := r.fs.Fetch(id)
		if err != nil {
			return "", false
		}

		r.c.Set(convertedVal, value, cache.DefaultExpiration)
		return value, true
	}
}
