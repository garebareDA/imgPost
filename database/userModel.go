package database

import(
	"github.com/jinzhu/gorm"
)

type UserData struct{
	gorm.Model
	ID string
	UserName string
	Icon string
}