package database

import(
	"github.com/jinzhu/gorm"
)

type UserData struct{
	gorm.Model
	userNmae string
	id string
	userImg string
}