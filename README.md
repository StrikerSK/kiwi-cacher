# Task description

Let us have a service that serves 2k requests per second. To serve each request, we need to access various data
provided by one of our many internal data sources. Obviously it's not feasible to fetch from these data sources with each
request that's fired our way, which means we need to cache the data sources somewhere.

You are provided with two different data sources that will serve as data providers for two separate cache instances.
The first source is an actual Kiwi.com API that provides location information and the second one is a SQLite database
provided in a separate file.

Your task is creating a cache, two data provider clients for the data sources defined below and writing a simple table test
to validate that the cache-provider orchestration works and returns correct data.

The cache must implement

```go
// Cacher defines methods used to obtain data from cache.
type Cacher interface {
	// Get returns id translated to string; boolean for cache hit/miss.
	Get(id int) (string, bool)
}
```

interface.

There is also [Fetcher](cache.go#L4) interface, which must be implemented by the two data provider clients you created.
The cache must **periodically** [FetchAll](cache.go#L9) fresh data from a data provider and, upon a cache miss,
it falls back to [Fetch](cache.go#L10) a single item from the data provider directly.

# Data sources

### Locations API [(docs)](https://docs.kiwi.com/locations/)
 * [GET](https://api.skypicker.com/locations?type=general&key=int_id&value=2210) a single location
 * [GET](https://api.skypicker.com/locations/graphql?query=%7Bdump%20%28options%3A%20%7Blocation_types%3A%20%5B%22airport%22%5D%2C%20active_only%3A%20%22true%22%7D%29%20%7Bint_id%20id%7D%7D) all locations

### SQLite database with ISO4217
* [Download](https://storage.googleapis.com/kw-gonuts/public/currencies.tar.gz) SQLite database
* `$ sqlite3 -column currencies.db "SELECT * from ISO4217;`
