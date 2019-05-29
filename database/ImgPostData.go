package database

import(
	"github.com/jinzhu/gorm"
)

type ImgPostData struct{
	gorm.Model
	PostID int
	UserID string
	UserName string
	Text string
}