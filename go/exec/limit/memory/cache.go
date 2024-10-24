package memory

import (
	"runtime"
	"sync"
	"time"
)

type CacheWrapper struct {
	*Cache
}

type cleaner struct {
	interval time.Duration
	stop     chan bool
}

func (cleaner *cleaner) Run(cache *Cache) {
	ticker := time.NewTicker(cleaner.interval)
	for {
		select {
		case <-ticker.C:
			cache.Clean()
		case <-cleaner.stop:
			ticker.Stop()
			return
		}
	}
}

func stopCleaner(wrapper *CacheWrapper) {
	wrapper.cleaner.stop <- true
	wrapper.cleaner = nil
}

func startCleaner(cache *Cache, interval time.Duration) {
	cleaner := &cleaner{
		interval: interval,
		stop:     make(chan bool),
	}
	cache.cleaner = cleaner
	go cleaner.Run(cache)
}

type Counter struct {
	mutex      sync.Mutex
	value      int64
	expiration int64
}

func (counter *Counter) Value() int64 {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.value
}

func (counter *Counter) Expiration() int64 {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.expiration
}

func (counter *Counter) Expired() bool {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.expiration == 0 || time.Now().UnixNano() > counter.expiration
}

func (counter *Counter) Load(expiration int64) (int64, int64) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	if counter.expiration == 0 || time.Now().UnixNano() > counter.expiration {
		return 0, expiration
	}
	return counter.value, counter.expiration
}

func (counter *Counter) Increment(value int64, expiration int64) (int64, int64) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	if counter.expiration == 0 || time.Now().UnixNano() > counter.expiration {
		counter.value = value
		counter.expiration = expiration
		return counter.value, counter.expiration
	}
	counter.value += value
	return counter.value, counter.expiration
}

type Cache struct {
	counters sync.Map
	cleaner  *cleaner
}

func NewCache(clearInterval time.Duration) *CacheWrapper {
	cache := &Cache{}
	wrapper := &CacheWrapper{Cache: cache}
	if clearInterval > 0 {
		startCleaner(cache, clearInterval)
		runtime.SetFinalizer(wrapper, stopCleaner)
	}
	return wrapper
}

func (cache *Cache) LoadOrStore(key string, counter *Counter) (*Counter, bool) {
	val, loaded := cache.counters.LoadOrStore(key, counter)
	if !loaded {
		return counter, false
	}
	actual := val.(*Counter)
	return actual, loaded
}

func (cache *Cache) Load(key string) (*Counter, bool) {
	val, ok := cache.counters.Load(key)
	if !ok || val == nil {
		return nil, false
	}

	actual := val.(*Counter)
	return actual, true
}

func (cache *Cache) Store(key string, counter *Counter) {
	cache.counters.Store(key, counter)
}

func (cache *Cache) Delete(key string) {
	cache.counters.Delete(key)
}

func (cache *Cache) Range(handler func(key string, counter *Counter)) {
	cache.counters.Range(func(k interface{}, v interface{}) bool {
		if v == nil {
			return true
		}

		key := k.(string)
		counter := v.(*Counter)
		handler(key, counter)
		return true
	})
}

func (cache *Cache) Increment(key string, value int64, duration time.Duration) (int64, time.Time) {
	expiration := time.Now().Add(duration).UnixNano()
	counter, loaded := cache.Load(key)
	if loaded {
		value, expiration := counter.Increment(value, expiration)
		return value, time.Unix(0, expiration)
	}

	counter, loade := cache.LoadOrStore(key, &Counter{
		mutex:      sync.Mutex{},
		value:      value,
		expiration: expiration,
	})

	if loade {
		value, expiration = counter.Increment(value, expiration)
		return value, time.Unix(0, expiration)
	}

	// Otherwise, it has been created, return given value.
	return value, time.Unix(0, expiration)

}

func (cache *Cache) Get(key string, duration time.Duration) (int64, time.Time) {
	expiration := time.Now().Add(duration).UnixNano()
	counter, ok := cache.Load(key)
	if !ok {
		return 0, time.Unix(0, expiration)
	}
	value, expiration := counter.Load(expiration)
	return value, time.Unix(0, expiration)
}

func (cache *Cache) Clean() {
	cache.Range(func(key string, counter *Counter) {
		if counter.Expired() {
			cache.Delete(key)
		}
	})
}

func (cache *Cache) Reset(key string, duration time.Duration) (int64, time.Time) {
	cache.Delete(key)
	expiration := time.Now().Add(duration).UnixNano()
	return 0, time.Unix(0, expiration)
}
