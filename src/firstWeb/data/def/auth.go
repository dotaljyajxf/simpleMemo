package def

import (
	"github.com/jinzhu/gorm"
)

//table
type Auth struct {
	gorm.Model
	Name     string `gorm:"unique_index;"`
	PassWord string `gorm:"type:varchar(32)"`
}
