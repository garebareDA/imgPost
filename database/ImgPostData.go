package database

import(
	"github.com/jinzhu/gorm"
)

type ImgPostData struct{
	gorm.Model
	userNmae string
	id string
	text string
	imgUrl string
}