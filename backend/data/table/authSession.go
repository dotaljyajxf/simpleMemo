package table

type TAuthSession struct {
	Uid      uint64 `sql:"uid"`
	NickName string `sql:"nick_name"`
	Account  string `sql:"account"`
	Mail     string `sql:"mail"`
	PhoneNum string `sql:"phone_num"`
}

/*
  TABLE_NAME auth
  PRIMARY KEY (`uid`),
  UNIQUE KEY `account` (`account`),
  UNIQUE KEY `nick_name` (`nick_name`),
  UNIQUE KEY `mail` (`mail`),
*/
