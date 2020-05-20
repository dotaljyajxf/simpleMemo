package auth

import (
	"firstWeb/data"
	"firstWeb/data/table"
	log "github.com/sirupsen/logrus"
)

func FindAuthObj(account string) (*table.Auth, error) {
	auth := table.NewAuth()
	auth.Account = account
	err := data.Db.Where(auth).First(auth).Error
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return nil, err
	}
	return auth, nil
}

func CreateAuth(name string, password string, mail, phoneNum, account string) (*table.Auth, error) {
	auth := table.NewAuth()
	auth.SetAccount(account)
	auth.SetNickName(name)
	auth.SetMail(mail)
	auth.SetPassWord(password)
	auth.SetPhoneNum(phoneNum)
	err := data.Db.Create(auth).Error
	if err != nil {
		log.Errorf("create auth error : %s", err.Error())
		return nil, err
	}
	return auth, nil
}
