package cache

import "time"

type keyPair struct {
	value    string
	deadline time.Time
}

type Cache struct {
	keyPairs map[string]keyPair
}

func NewCache() Cache {
	return Cache{}
}

func (r *Cache) Get(key string) (string, bool) {
	cache, ok := r.keyPairs[key]
	currentTime := time.Now()

	if currentTime.After(cache.deadline) {
		delete(r.keyPairs, key)

		return "", false
	}

	if !ok {
		return "", false
	}

	return cache.value, true
}

func (r *Cache) Put(key, value string) {
	r.keyPairs[key] = keyPair{
		value:    value,
		deadline: time.Time{},
	}
}

func (r *Cache) Keys() []string {
	keys := make([]string, 0, len(r.keyPairs))
	currentTime := time.Now()

	for key, value := range r.keyPairs {
		if currentTime.After(value.deadline) {
			delete(r.keyPairs, key)

			continue
		}

		keys = append(keys, key)
	}

	return keys
}

func (r *Cache) PutTill(key, value string, deadline time.Time) {
	r.keyPairs[key] = keyPair{
		value:    value,
		deadline: deadline,
	}
}
