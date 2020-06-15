package util

import (
	"firstWeb/conf"
	"firstWeb/proto/pb"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(conf.Config.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		sess := sessions.Default(c)
		auth := sess.Get("user").(pb.TAuthInfo)
		token := auth.GetToken()
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			code = -1
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = -2
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = -3
			}
		}

		retAuth := pb.NewTAuthInfo()
		defer retAuth.Put()

		if code != 0 {
			ret := pb.NewTAppRet()
			MakeErrRet(ret, http.StatusForbidden, "AuthError")
			c.ProtoBuf(ret.Code, ret)
			c.Abort()
			return
		}

		c.Next()
	}
}
