package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"github.com/olahol/go-imageupload"
	"net/http"
	"imgPost/database"
)

func UserSetting(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	user := c.Param("user")

	token := csrf.GetToken(c)
	isAlive(alive, c)

	if userID != user {
		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()
	}

	c.HTML(http.StatusOK, "acount.html", gin.H{
		"_csrf": token,
	})
}

func UserSettingPost(c *gin.Context){
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	user := c.Param("user")

	isAlive(alive, c)

	if userID == user {
		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()
	}

	c.Request.ParseForm()
	name := c.Request.Form["name"]

	imageupload.LimitFileSize(5242880, c.Writer, c.Request)

	//アップロード画像の取得
	var icon string
	img, err := imageupload.Process(c.Request, "file")
	if err == nil {
		icon = userID.(string)
		img.Save("./icon/" + userID.(string) + ".png")
	}else{
		icon = "NoIcon"
	}

	db := database.ConnectDB()
	defer db.Close()

	userData := database.UserData{}
	userDataUpdate := database.UserData{}
	userData.UserID = userID.(string)

	db.First(&userData)

	if name == nil {
		userDataUpdate.UserName = userData.UserName
	}else{
		userDataUpdate.UserName = name[1]
	}

	userDataUpdate.Icon = icon

	db.Save(&userDataUpdate)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}