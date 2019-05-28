package database

import(
	"github.com/jinzhu/gorm"
)

type ImgPostData struct{
	gorm.Model
	UserID string
	UserName string
	Text string
	ImgURL int
	PostID int
}