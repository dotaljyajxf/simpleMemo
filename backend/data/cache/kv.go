package cache

import (
	"backend/conf"
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Kv *KvCache

type KvCache struct {
	conn_pool *redis.Pool
}

func InitKvCache() {
	Kv = new(KvCache)
	Kv.initPool()
}

func (c *KvCache) initPool() {
	c.conn_pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf(conf.Config.KvRedisHost),
				redis.DialPassword(conf.Config.KvRedisPassWd),
				redis.DialDatabase(conf.Config.KvRedisDB),
				redis.DialConnectTimeout(time.Second*2),
				redis.DialReadTimeout(time.Second*2),
				redis.DialWriteTimeout(time.Second*2),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     conf.Config.KvRedisMaxIdel,
		MaxActive:   conf.Config.KvRedisMaxActive,
		IdleTimeout: time.Second * time.Duration(conf.Config.CacheRedisIdelTimeout),
		Wait:        true,
	}

	if _, err := c.Do("PING"); err != nil {
		panic(err)
	}
}

func (c *KvCache) Do(command string, args ...interface{}) (reply interface{}, err error) {
	if c.conn_pool == nil {
		return nil, errors.New("NotInitRedis")
	}
	conn := c.conn_pool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}
