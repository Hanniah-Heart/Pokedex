package pokecache
import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mutex *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	var newCache Cache
	newCache.reapLoop(interval)
	return &newCache
}

func (cache Cache) Add(key string, val []byte) {
	_, ok := cache.entry[key]
	if ok {
		return
	}
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	newEntry := cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
	cache.entry[key] = newEntry
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	targetEntry, ok := cache.entry[key]
	if !ok {
		var nullVal []byte
		return nullVal, ok
	}
	return targetEntry.val, true
}

func (cache Cache) reapLoop(interval time.Duration) {
	// after interval seconds, destroy associated cache
	// should use a sync.mutex
}

/*

5. Update your code that makes requests to the PokeAPI to use the cache. Create the cache once and reuse it in your PokeAPI request layer. If you already have the data for a given URL (which is our cache key) in the cache, you should use that instead of making a new request. Whenever you do make a request, you should add the response to the cache.
6. Write at least 1 test for your cache package! The tip below should help you get started.
7. Test your application manually to make sure that the cache works as expected. When you use the map command to get data for the first time there should be a noticeable waiting time. However, when you use mapb it should be instantaneous because the data for that page is already in the cache. Feel free to add some logging that informs you in the command line when the cache is being used.

Tips
Clearing the Cache
You can use a time.Ticker inside a goroutine started by NewCache. In a loop like for range ticker.C { ... }, check the entries and remove any whose createdAt is older than the cache's interval.

Running Tests
You can run tests for all packages in a Go module by running go test ./... from the root of the module.

*/



