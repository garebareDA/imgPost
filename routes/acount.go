package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/olahol/go-imageupload"
	"github.com/utrack/gin-csrf"
	"net/http"
	"log"
	"imgPost/database"
)

//CreateAcount アカウント作成
func CreateAcount(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	isAlive(alive, c)

	c.Request.ParseForm()
	name := c.Request.Form["name"]

	imageupload.LimitFileSize(5242880, c.Writer, c.Request)

	//アップロード画像の取得
	var icon string
	img, err := imageupload.Process(c.Request, "file")
	if err == nil {
		icon = userID.(string)
		img.Save("./icon/" + userID.(string) + ".jpg")
	}else{
		icon = "NoIcon"
	}

	db := database.ConnectDB()
	defer db.Close()

	db.Where(database.UserData{UserID: userID.(string)}).Assign(database.UserData{UserName: name[0], Icon: icon}).FirstOrCreate(&database.UserData{})
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

//Acount アカウントの作成画面
func Acount(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	isAlive(alive, c)

	token := csrf.GetToken(c)

	db := database.ConnectDB()
	defer db.Close()

	user := database.UserData{}
	user.UserID = userID.(string)

	if db.First(&user).RecordNotFound() == false {
		log.Println(db.First(&user))
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	c.HTML(http.StatusOK, "acount.html", gin.H{
			"_csrf": token,
			"alive": alive,
			"id":userID,
		})
}