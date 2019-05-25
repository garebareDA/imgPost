package database

import(
	"github.com/jinzhu/gorm"
)

type ImgPostData struct{
	gorm.Model
	UserID string
	UserNmae string
	Text string
	ImgURL string
	postID string
}