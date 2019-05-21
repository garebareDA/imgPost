package main

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/olahol/go-imageupload"
	"net/http"
	"log"
	"imgPost/database"
)

func home(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")
	log.Println(userID)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"h": "<h1>aaaaaaaaaaaaaaaaa</h1>",
		"alive": alive,
		"id":userID,
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

	//TODO uuidで管理
	thumb.Save("./icon" + "uuid")

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

	db.Where(database.UserData{ID: userID.(string)}).Assign(database.UserData{UserName: name[0], Icon:""}).FirstOrInit(database.UserData{})
}