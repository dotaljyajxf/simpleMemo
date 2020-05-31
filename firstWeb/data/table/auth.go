package table

type Auth struct {
	Uid        uint64 `gorm:"column:uid;primary_key;AUTO_INCREMENT"`
	NickName   string `gorm:"unique_index;"`
	Account    string `gorm:"unique_index;"`
	Mail       string
	PassWord   string `gorm:"type:varchar(32)"`
	PhoneNum   string
	CreateTime int64
}
