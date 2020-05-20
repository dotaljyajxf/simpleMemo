package api

import (
	"encoding/json"
	"firstWeb/module/auth"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Engine) {
	c.GET("/Login", func(c *gin.Context) {
		account := c.Param("account")
		password := c.Param("password")

		authObj, err := auth.FindAuthObj(account)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "authError"})
			return
		}

		if authObj.GetPassWord() != password {
			c.JSON(http.StatusOK, gin.H{"message": "passwordErr"})
			return
		}

		token, err := util.GenerateToken(authObj.GetAccount(), authObj.GetPassWord())

		if err != nil {
			//...
		}
		retStr, _ := json.Marshal(authObj)
		authObj.Release()
		c.JSON(http.StatusOK, gin.H{"authObj": retStr, "token": token})
	})
}

func Regist(c *gin.Engine) {
	c.GET("/Regist", func(c *gin.Context) {
		account := c.Param("account")
		passWord := c.Param("password")
		mail := c.Param("mail")
		phoneNum := c.Param("phonenum")
		nickName := c.Param("name")

		if account == "" || passWord == "" {
			c.JSON(http.StatusOK, gin.H{"message": "must need account and password"})
			return
		}
		authObj, err := auth.CreateAuth(nickName, passWord, mail, phoneNum, account)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}

		token, err := util.GenerateToken(authObj.GetAccount(), authObj.GetPassWord())

		if err != nil {
			//...
		}
		retStr, _ := json.Marshal(authObj)
		authObj.Release()
		c.JSON(http.StatusOK, gin.H{"authObj": retStr, "token": token})
	})
}
