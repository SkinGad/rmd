package cache

func NewCache() (*Cache, error) {
	cache := Cache{}
	err := cache.Load()
	return &cache, err
}
