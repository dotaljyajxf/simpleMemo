package api

import (
	"backend/model/auth"
	"backend/proto/pb"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Login(c *gin.Context, ret *pb.TAppRet) error {
	account := c.PostForm("account")
	password := c.PostForm("password")

	retAuth := pb.NewTAuthInfo()
	authObj, err := auth.FindAuthObj(account)
	defer authObj.Put()
	if err != nil {
		return util.MakeErrRet(ret, http.StatusOK, err.Error())
	}
	if authObj.PassWord != password {
		return util.MakeErrRet(ret, http.StatusOK, "PassWordError")
	}

	token, err := util.GenerateToken(authObj.Account, authObj.PassWord)

	if err != nil {
		return util.MakeErrRet(ret, http.StatusOK, "tokenError")
	}

	retAuth.PhoneNum = authObj.PhoneNum
	retAuth.Mail = authObj.Mail
	retAuth.Token = token
	retAuth.Uid = authObj.Uid
	retAuth.NickName = authObj.NickName

	sessionUser := new(auth.SessionUser)
	sessionUser.Account = authObj.Account
	sessionUser.PhoneNum = authObj.PhoneNum
	sessionUser.Mail = authObj.Mail
	sessionUser.NickName = authObj.NickName
	sessionUser.Uid = authObj.Uid
	sessionUser.Token = token
	err = auth.SetAuthSession(token, sessionUser)
	if err != nil {
		log.Errorf("set session error %s", err.Error())
	}

	c.SetCookie("token", token, 300, "/", "127.0.0.1", false, true)
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
	defer authObj.Put()
	if err != nil {
		log.Infof("create auth failed ret:%s", err.Error())
		return util.MakeErrRet(ret, http.StatusOK, "CreateFaild")
	}

	token, err := util.GenerateToken(authObj.Account, authObj.PassWord)

	if err != nil {
		//...
	}

	c.SetCookie("token", token, 300, "/", "127.0.0.1", false, true)

	retAuth.PhoneNum = authObj.PhoneNum
	retAuth.Mail = authObj.Mail
	retAuth.Token = token
	retAuth.Uid = authObj.Uid
	retAuth.NickName = authObj.NickName

	sessionUser := new(auth.SessionUser)
	sessionUser.Account = authObj.Account
	sessionUser.PhoneNum = authObj.PhoneNum
	sessionUser.Mail = authObj.Mail
	sessionUser.NickName = authObj.NickName
	sessionUser.Uid = authObj.Uid
	sessionUser.Token = token
	auth.SetAuthSession(token, sessionUser)

	return util.MakeSuccessRet(ret, http.StatusOK, retAuth)
}
