package cache

import (
	"errors"
	"firstWeb/conf"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var _conn_pool *redis.Pool

func InitCache() {
	_conn_pool = &redis.Pool{
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

	if _, err := Do("PING"); err != nil {
		panic(err)
	}
}

func Do(command string, args ...interface{}) (reply interface{}, err error) {
	if _conn_pool == nil {
		return nil, errors.New("NotInitRedis")
	}
	conn := _conn_pool.Get()
	defer conn.Close()
	return conn.Do(command, args)
}

func Get(key string) (reply interface{}, err error) {
	return Do("GET", key)
}

func Set(key string, val interface{}) (reply interface{}, err error) {
	return Do("SET", key, val)
}

func Del(key string) (reply interface{}, err error) {
	return Do("Del", key)
}
