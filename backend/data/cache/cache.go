package cache

type CacheHandler interface {
	Get(key string) (reply interface{}, err error)
	Set(key string, val interface{}) (reply interface{}, err error)
	Del(key string) (reply interface{}, err error)
}

type Cache struct {
	c CacheHandler
}

func (c Cache) Get(key string) (reply interface{}, err error) {
	return c.do("GET", key)
}

func (c Cache) Set(key string, val interface{}) (reply interface{}, err error) {
	return c.do("SET", key, val)
}

func (c Cache) Del(key string) (reply interface{}, err error) {
	return c.do("Del", key)
}
