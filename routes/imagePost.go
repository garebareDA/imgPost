package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/olahol/go-imageupload"
	"net/http"
	"log"
	"strconv"
	"imgPost/database"
)

//ImagePost アップロード画像の保存
func ImagePost(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	isAlive(alive.(bool), c)

	c.Request.ParseForm()
	text := c.Request.Form["text"]

	if text == nil {
		text[0] = "(テキストはありません)"
	}

	imageupload.LimitFileSize(5242880, c.Writer, c.Request)

	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		log.Panicln(err)
	}

	db := database.ConnectDB()
	defer db.Close()



	userData := database.UserData{}
	userData.UserID = userID.(string)
	user := db.First(&userData)
	log.Println(user)

	if user.RecordNotFound() == true {
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	lastpost := database.ImgPostData{}
	db.Last(&lastpost)
	lastID := lastpost.PostID

	var num int

	if lastID == 0 {
		num = 1
	}else{
		num = lastID + 1
	}

	log.Println(num)

	postData := database.ImgPostData{}
	postData.UserID = userID.(string)
	postData.UserName = userData.UserName
	postData.Text = text[0]
	postData.PostID = num

	db.Create(&postData)

	img.Save("./img/" + strconv.Itoa(num) + ".jpg")

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}