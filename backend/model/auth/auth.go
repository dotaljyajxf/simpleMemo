package auth

import (
	"backend/data"
	"backend/data/table"
	"context"

	log "github.com/sirupsen/logrus"
)

func FindAuthObj(account string) (*table.TAuth, error) {
	auth := table.NewTAuth()
	log.Info("FindAuthObj Begin")
	err := data.Manager.Query(auth, auth.SelectByAccountSql(), account)
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return nil, err
	}
	log.Info("FindAuthObj end")
	return auth, nil
}

func CreateAuth(name string, password string, mail, phoneNum, account string) (*table.TAuth, error) {
	auth := table.NewTAuth()
	auth.Account = account
	auth.NickName = name
	auth.Mail = mail
	auth.PassWord = password
	auth.PhoneNum = phoneNum

	_, err := data.Manager.InsertTable(context.Background(), auth)
	if err != nil {
		log.Errorf("create auth error : %s", err.Error())
		return nil, err
	}
	return auth, nil
}
