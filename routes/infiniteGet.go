package routes

import(
	"github.com/gin-gonic/gin"
	"imgPost/database"
	"strconv"
)

type ImgPostDataJson struct{
	ID int `json:"id"`
	UserID string `json:"userID"`
	UserName string `json:"userName"`
	Text string `josn:"text"`
	Icon string `json:"icon"`
}

func InfiniteGet(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()

	page := c.Query("page")
	last := c.Query("last")

	p, err := strconv.Atoi(page)
	if err != nil{
		panic(err)
	}

	l, err := strconv.Atoi(last)
	if err != nil{
		panic(err)
	}

	imageJson := []ImgPostDataJson{}
	imgPostt := database.ImgPostData{}
	userData := database.UserData{}

	db.Last(&imgPostt)
	lastID := imgPostt.PostID - p

	for {
		imgPost := database.ImgPostData{}
		imgPost.PostID = lastID
		db.Where("post_id = ?", lastID).First(&imgPost)

		if db.Where("post_id = ?", lastID).First(&imgPost).RecordNotFound() == true {
			lastID--
			l--
			continue
		}

		userData.UserID = imgPost.UserID
		db.First(&userData)

		data := ImgPostDataJson{
			ID: imgPost.PostID,
			UserID: imgPost.UserID,
			UserName: imgPost.UserName,
			Text: imgPost.Text,
			Icon: userData.Icon,
		}

		imageJson = append(imageJson, data)

		lastID--
		l--

		if l <= 1 || lastID <= 0 {
			break
		}else{
			continue
		}

	}

	c.JSON(200, imageJson)
}