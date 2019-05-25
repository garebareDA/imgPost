package database

import(
	"github.com/jinzhu/gorm"
)

type UserData struct{
	gorm.Model
	UserID string
	UserName string
	Icon string
}