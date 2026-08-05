package pokecache
import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt time.Time
	Val []byte
}

type Cache struct {
	Entry map[string]cacheEntry
	mutex *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache { // consider changing this later to allow for runtime mutable configuration of the interval so that the user can adjust their memory usage to their preference during play in case we don't want them to require restarting.
	var newCache Cache
	go newCache.reapLoop(interval)
	return &newCache
}

func (cache Cache) Add(key string, val []byte) {
	_, ok := cache.Entry[key]
	if ok {
		return
	}
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	newEntry := cacheEntry {
		createdAt: time.Now(),
		Val: val,
	}
	cache.Entry[key] = newEntry
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	targetEntry, ok := cache.Entry[key]
	if !ok {
		var nullVal []byte
		return nullVal, ok
	}
	//targetEntry.createdAt = time.Now() //unrequested feature to reset the reaping clock when the cache entry is gotten
	return targetEntry.Val, true
}

func (cache Cache) reapLoop(interval time.Duration) { // consider changing this later to allow for runtime mutable configuration of the interval so that the user can adjust their memory usage to their preference during play in case we don't want them to require restarting.
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for true {
		<- ticker.C
		for key, _ := range cache.Entry {
			destroyAt := cache.Entry[key].createdAt.Add(interval)
			if time.Now().After(destroyAt) {
				cache.mutex.Lock()
				delete(cache.Entry, key)
				cache.mutex.Unlock()
			}
		}
	}
}

/*

6. Write at least 1 test for your cache package! The tip below should help you get started.
7. Test your application manually to make sure that the cache works as expected. When you use the map command to get data for the first time there should be a noticeable waiting time. However, when you use mapb it should be instantaneous because the data for that page is already in the cache. Feel free to add some logging that informs you in the command line when the cache is being used.

Tips
Clearing the Cache
You can use a time.Ticker inside a goroutine started by NewCache. In a loop like for range ticker.C { ... }, check the entries and remove any whose createdAt is older than the cache's interval.

Running Tests
You can run tests for all packages in a Go module by running go test ./... from the root of the module.

*/



