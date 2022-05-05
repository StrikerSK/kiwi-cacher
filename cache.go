package cache

// Fetcher defines methods used to obtain data from provider.
type Fetcher interface {
	// Fetch obtains string representation of id from provider.
	// It should be used as fallback if cache does not already contain the id.
	Fetch(id int) (string, error)
	// FetchAll data from provider.
	FetchAll() (map[int]string, error)
}

type Cache struct{}

func NewCache(fs Fetcher) *Cache {
	panic("implement me")
}
