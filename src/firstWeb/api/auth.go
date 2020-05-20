package api

import (
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

		retAuth := pb.NewTAuthInfo()
		defer retAuth.Put()
		authObj, err := auth.FindAuthObj(account)
		if err != nil {
			retAuth.SetMessage("AccountError")
			c.ProtoBuf(http.StatusOK, retAuth)
			return
		}

		if authObj.GetPassWord() != password {
			retAuth.SetMessage("PassWordError")
			c.ProtoBuf(http.StatusOK, retAuth)
			return
		}

		token, err := util.GenerateToken(authObj.GetAccount(), authObj.GetPassWord())

		if err != nil {
			//...
		}

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

		retAuth := pb.NewTAuthInfo()
		defer retAuth.Put()

		log.Infof("regist: %s,%s", account, passWord)
		if account == "" || passWord == "" {
			retAuth.SetMessage("NeedAccountAndPassWord")
			c.ProtoBuf(http.StatusOK, retAuth)
			return
		}

		authObj, err := auth.CreateAuth(nickName, passWord, mail, phoneNum, account)
		if err != nil {
			retAuth.SetMessage(err.Error())
			c.ProtoBuf(http.StatusOK, retAuth)
			return
		}

		token, err := util.GenerateToken(authObj.GetAccount(), authObj.GetPassWord())

		if err != nil {
			//...
		}

		retAuth.SetPhoneNum(authObj.GetPhoneNum())
		retAuth.SetMail(authObj.GetMail())
		retAuth.SetToken(token)
		retAuth.SetUid(string(authObj.GetUid()))
		retAuth.SetNickName(authObj.GetAccount())

		c.ProtoBuf(http.StatusOK, retAuth)
		authObj.Release()

		//retStr, _ := json.Marshal(authObj)
		//authObj.Release()
		//c.JSON(http.StatusOK, gin.H{"authObj": retStr, "token": token})
	})
}
