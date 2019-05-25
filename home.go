package main

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/olahol/go-imageupload"
	"github.com/utrack/gin-csrf"
	"github.com/google/uuid"
	"net/http"
	"log"
	"imgPost/database"
)

func home(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	token := csrf.GetToken(c)

	log.Println(userID)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"h": "<h1>aaaaaaaaaaaaaaaaa</h1>",
		"alive": alive,
		"id":userID,
		"_csrf": token,
	})
}

//Registration アカウント作成
func Registration(c *gin.Context) {
	//アップロード画像の取得
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		log.Panicln(err)
	}

	//300x300にリサイズ
	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)
	if err != nil {
		log.Panicln(err)
	}

	uuid := uuid.New().String()

	thumb.Save("./icon/" + uuid + ".png")

	c.Request.ParseForm()
	name := c.Request.Form["name"]

	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	if alive == false {
		session.Clear()
		session.Save()
		c.Redirect(http.StatusOK, "/")
	}

	db := database.ConnectDB()
	defer db.Close()

	db.Where(database.UserData{UserID: userID.(string)}).Assign(database.UserData{UserName: name[0], Icon: uuid}).FirstOrCreate(&database.UserData{})
	c.Redirect(http.StatusMovedPermanently, "/")
}

func acount(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	token := csrf.GetToken(c)

	c.HTML(http.StatusOK, "acount.html", gin.H{
		"_csrf": token,
		"alive": alive,
		"id":userID,
	})
}