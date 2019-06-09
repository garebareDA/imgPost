package routes

import(
	"github.com/gin-gonic/gin"
	"imgPost/database"
	"strconv"
	"github.com/gin-contrib/sessions"
	"log"
)

type ImgPostDataJson struct{
	PostID int `json:"id"`
	UserID string `json:"userID"`
	UserName string `json:"userName"`
	Text string `josn:"text"`
	Icon string `json:"icon"`
	Delete bool `json:"delete"`
}

func InfiniteGet(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()

	page := c.Query("page")
	last := c.Query("last")


	p, err := strconv.Atoi(page)
	if err != nil{
		c.String(400,"error")
	}

	l, err := strconv.Atoi(last)
	if err != nil{
		c.String(400,"error")
	}

	imageJson := []ImgPostDataJson{}
	lastPost := database.ImgPostData{}
	userData := database.UserData{}

	db.Last(&lastPost)
	lastID := lastPost.PostID - p

	for {
		imgPost := database.ImgPostData{}
		db.Where("post_id = ?", lastID).First(&imgPost)

		if db.Where("post_id = ?", lastID).First(&imgPost).RecordNotFound() == true {

			lastID--
			l--

			if l <= 1 || lastID <= 0 {
				break
			}else{
				continue
			}
		}

		db.Where("user_id = ?", imgPost.UserID).First(&userData)

		log.Println(userData.Icon)

		data := ImgPostDataJson{
			PostID: imgPost.PostID,
			UserID: imgPost.UserID,
			UserName: imgPost.UserName,
			Text: imgPost.Text,
			Icon: imgPost.UserID,
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

func InfiniteGetUser(c *gin.Context) {

	db := database.ConnectDB()
	defer db.Close()

	userID := c.Query("id")
	last := c.Query("last")
	l, _:= strconv.Atoi(last)


	imageJson := []ImgPostDataJson{}
	lastPost := []database.ImgPostData{}

	db.Where("user_id = ?", userID).Find(&lastPost)

	lastnum := len(lastPost) - l

	log.Println(lastnum)

	session := sessions.Default(c)
	userId := session.Get("userId")
	alive := session.Get("alive")
	delete := false

	if alive != nil && userId.(string) == userID && alive.(bool) == true{
		delete = true
	}

	for lastnum > 0 {
		lastnum --
		post := lastPost[lastnum]
		userData := database.UserData{}
		userData.UserID = post.UserID
		db.Where("user_id = ?", post.PostID).First(&userData)

		data := ImgPostDataJson{
			PostID: post.PostID,
			UserID: post.UserID,
			UserName: post.UserName,
			Text: post.Text,
			Icon: userData.UserID,
			Delete: delete,
		}

		imageJson = append(imageJson, data)
	}

	c.JSON(200, imageJson)
}