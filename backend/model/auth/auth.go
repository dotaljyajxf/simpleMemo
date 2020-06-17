package auth

import (
	"backend/data"
	"backend/data/table"
	"context"

	log "github.com/sirupsen/logrus"
)

func FindAuthObj(account string) (*table.TAuth, error) {
	auth := table.NewTAuth()
	auth.Account = account
	err := data.Db.Where(auth).First(auth).Error
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return nil, err
	}
	return auth, nil
}

func CreateAuth(name string, password string, mail, phoneNum, account string) (*table.TAuth, error) {
	auth := table.NewTAuth()
	auth.SetAccount(account)
	auth.SetNickName(name)
	auth.SetMail(mail)
	auth.SetPassWord(password)
	auth.SetPhoneNum(phoneNum)

	_, err := data.Data.InsertTable(context.Background(), auth)
	if err != nil {
		log.Errorf("create auth error : %s", err.Error())
		return nil, err
	}
	return auth, nil
}
