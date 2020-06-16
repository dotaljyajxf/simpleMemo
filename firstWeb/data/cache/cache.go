package cache

import (
	"errors"
	"firstWeb/conf"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Cache struct {
	conn_pool *redis.Pool
}


func NewRedisCache() *Cache {
	cache := new(Cache)
	cache.initCache()
	return cache
}

func (c *cache)initCache() {
	c.conn_pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf(conf.Config.RedisHost),
				//redis.DialPassword(cfg.Password),
				redis.DialDatabase(conf.Config.RedisDB),
				redis.DialConnectTimeout(time.Second*2),
				redis.DialReadTimeout(time.Second*2),
				redis.DialWriteTimeout(time.Second*2),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     conf.Config.RedisMaxIdel,
		MaxActive:   conf.Config.RedisMaxActive,
		IdleTimeout: time.Duration(conf.Config.RedisIdelTimeout), // TODO 调整为稍大点的值?
		Wait:        true,
	}

	if _, err := c.do("PING"); err != nil {
		panic(err)
	}
}

func (c *cache)do(command string, args ...interface{}) (reply interface{}, err error) {
	if c.conn_pool == nil {
		return nil, errors.New("NotInitRedis")
	}
	conn := c.conn_pool.Get()
	defer conn.Close()
	return conn.Do(command, args)
}

func (c *cache)Get(key string) (reply interface{}, err error) {
	return c.do("GET", key)
}

func (c *cache)Set(key string, val interface{}) (reply interface{}, err error) {
	return c.do("SET", key, val)
}

func (c *cache)Del(key string) (reply interface{}, err error) {
	return c.do("Del", key)
}
