package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

type SessionUser struct {
	Uid      int64
	NickName string
	Account  string
	Mail     string
	PhoneNum string
	Token    string
}

func GetAuthSession(c *gin.Context) *SessionUser {
	sess := sessions.Default(c)
	return sess.Get("user").(*SessionUser)
}

func SetAuthSession(c *gin.Context, s *SessionUser) error {
	sess := sessions.Default(c)
	sess.Set("user", s)
	return sess.Save()
}
