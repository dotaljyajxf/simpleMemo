package api

import (
	"encoding/json"
	"firstWeb/module/auth"
	"firstWeb/proto/pb"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Login(c *gin.Engine) {
	c.POST("/Login", func(c *gin.Context) {
		account := c.PostForm("account")
		password := c.PostForm("password")

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

		retAuth := pb.NewTAuthInfo()
		defer retAuth.Put()
		retAuth.SetPhoneNum(authObj.GetPhoneNum())
		retAuth.SetMail(authObj.GetMail())
		retAuth.SetToken(token)
		retAuth.SetUid(string(authObj.GetUid()))
		retAuth.SetNickName(authObj.GetAccount())

		//retStr, _ := json.Marshal(authObj)
		//log.Infof("retStr: %s\n", retStr)

		c.ProtoBuf(http.StatusOK, retAuth)
		authObj.Release()
	})
}

func Regist(c *gin.Engine) {
	c.POST("/Regist", func(c *gin.Context) {
		account := c.PostForm("account")
		passWord := c.PostForm("password")
		mail := c.PostForm("mail")
		phoneNum := c.PostForm("phonenum")
		nickName := c.PostForm("name")

		log.Infof("regist: %s,%s", account, passWord)
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
