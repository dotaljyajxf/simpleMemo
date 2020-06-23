package auth

import (
	"backend/data/cache"
	"encoding/json"

	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SessionUser struct {
	Uid      int64  `json:"uid"`
	NickName string `json:"nick_name"`
	Account  string `json:"account"`
	Mail     string `json:"mail"`
	PhoneNum string `json:"phone_num"`
	Token    string `json:"token"`
}

const KEY_PREFIX = "game_"

func GetAuthSession(c *gin.Context) *SessionUser {
	key, _ := c.Cookie("token")
	key = KEY_PREFIX + key
	res, err := redis.Bytes(cache.Kv.Do("GET", key))
	if err != nil {
		logrus.Errorf("get session from cache error : %s,%s", err.Error(), key)
		return nil
	}
	sess := new(SessionUser)
	err = json.Unmarshal(res, sess)
	if err != nil {
		logrus.Errorf("Unmarshal session from cache error : %s,%s", err.Error(), key)
		return nil
	}
	return sess
}

func SetAuthSession(key string, s *SessionUser) error {
	sJson, err := json.Marshal(s)
	if err != nil {
		logrus.Errorf("json session error : %s", err.Error())
		return err
	}

	key = KEY_PREFIX + key
	_, err = cache.Kv.Do("SETEX", key, 3600, sJson)
	if err != nil {
		logrus.Errorf("cache session error : %s,%s", err.Error(), key)
		return err
	}
	return nil
}
