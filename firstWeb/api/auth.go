package api

import (
	"firstWeb/module/auth"
	"firstWeb/proto/pb"
	"firstWeb/util"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context, ret *pb.TAppRet) error {
	account := c.PostForm("account")
	password := c.PostForm("password")

	retAuth := pb.NewTAuthInfo()
	defer retAuth.Put()
	authObj, err := auth.FindAuthObj(account)
	defer authObj.Release()
	if err != nil {
		return util.MakeErrRet(ret, http.StatusOK, err.Error())
	}

	if authObj.GetPassWord() != password {
		return util.MakeErrRet(ret, http.StatusOK, "PassWordError")
	}

	token, err := util.GenerateToken(authObj.GetAccount(), authObj.GetPassWord())

	if err != nil {
		//..
	}

	retAuth.PhoneNum = &authObj.PhoneNum
	retAuth.Mail = &authObj.Mail
	retAuth.Token = &token
	retAuth.Uid = &authObj.Uid
	retAuth.NickName = &authObj.NickName

	sess := sessions.Default(c)
	sess.Set("user", retAuth)
	err = sess.Save()

	return util.MakeSuccessRet(ret, http.StatusOK, retAuth)
}

func Regist(c *gin.Context, ret *pb.TAppRet) error {
	account := c.PostForm("account")
	passWord := c.PostForm("password")
	mail := c.PostForm("mail")
	phoneNum := c.PostForm("phonenum")
	nickName := c.PostForm("name")

	retAuth := pb.NewTAuthInfo()

	log.Infof("regist: %s,%s", account, passWord)
	if account == "" || passWord == "" {
		return util.MakeErrRet(ret, http.StatusOK, "NeedAccountAndPassWord")
	}

	authObj, err := auth.CreateAuth(nickName, passWord, mail, phoneNum, account)
	defer authObj.Release()
	if err != nil {
		log.Infof("create auth failed ret:%s", err.Error())
		return util.MakeErrRet(ret, http.StatusOK, "CreateFaild")
	}

	token, err := util.GenerateToken(authObj.GetAccount(), authObj.GetPassWord())

	if err != nil {
		//...
	}

	retAuth.PhoneNum = &authObj.PhoneNum
	retAuth.Mail = &authObj.Mail
	retAuth.Token = &token
	retAuth.Uid = &authObj.Uid
	retAuth.NickName = &authObj.NickName

	sess := sessions.Default(c)
	sess.Set("user", retAuth)
	err = sess.Save()

	return util.MakeSuccessRet(ret, http.StatusOK, retAuth)
}
