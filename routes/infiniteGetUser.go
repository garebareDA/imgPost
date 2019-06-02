package routes

import(
	"github.com/gin-gonic/gin"
	"imgPost/database"
	"strconv"
)

func InfiniteGetUser(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()

	userID := c.Query("id")
	last := c.Query("last")
	l, _:= strconv.Atoi(last)


	imageJson := []ImgPostDataJson{}
	lastPost := []database.ImgPostData{}

	db.Where("user_id = ?", userID).Find(&lastPost)

	lastnum := len(lastPost) - 1 - l

	for lastnum > 0 {
		post := lastPost[lastnum]
		userData := database.UserData{}
		userData.UserID = post.UserID
		db.First(&userData)

		data := ImgPostDataJson{
			PostID: post.PostID,
			UserID: post.UserID,
			UserName: post.UserName,
			Text: post.Text,
			Icon: userData.Icon,
		}

		imageJson = append(imageJson, data)

		lastnum --
	}

	c.JSON(200, imageJson)
}