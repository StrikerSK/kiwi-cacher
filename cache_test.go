package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockFetcher struct{}

func (r *MockFetcher) Fetch(id int) (string, error) {
	switch id {
	case 1:
		return "London", nil
	case 2:
		return "Berlin", nil
	case 3:
		return "Prague", nil
	default:
		return "", fmt.Errorf("location [%d] not found", id)
	}
}

func (r *MockFetcher) FetchAll() (map[int]string, error) {
	return map[int]string{
		1: "London",
		2: "Berlin",
		3: "Prague",
	}, nil
}

var testCache = NewCache(&MockFetcher{})

func TestItemFound(t *testing.T) {
	value, found := testCache.Get(1)

	assert.True(t, found)
	assert.Equal(t, "London", value)
}

func TestFollowingRuns(t *testing.T) {
	start := time.Now()
	_, _ = testCache.Get(1)
	firstRun := time.Since(start).Nanoseconds()

	start = time.Now()
	_, _ = testCache.Get(1)
	secondRun := time.Since(start).Nanoseconds()

	fmt.Printf("First time run: %d and second time run: %d\n", firstRun, secondRun)
	assert.Less(t, secondRun, firstRun)
}
