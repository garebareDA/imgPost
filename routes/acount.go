package routes

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

//CreateAcount アカウント作成
func CreateAcount(c *gin.Context) {

	imageupload.LimitFileSize(5242880, c.Writer, c.Request)

	//アップロード画像の取得
	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		log.Panicln(err)
	}

	log.Println(img.Size)

	uuid := uuid.New().String()

	img.Save("./icon/" + uuid + ".png")

	c.Request.ParseForm()
	name := c.Request.Form["name"]

	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	if alive == nil {
		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()
	}

	db := database.ConnectDB()
	defer db.Close()

	db.Where(database.UserData{UserID: userID.(string)}).Assign(database.UserData{UserName: name[0], Icon: uuid}).FirstOrCreate(&database.UserData{})
	c.Redirect(http.StatusFound, "/")
}

func Acount(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	token := csrf.GetToken(c)

	db := database.ConnectDB()
	defer db.Close()

	if alive == true {
		user := database.UserData{}
		user.UserID = userID.(string)

		log.Println(db.First(&user).RecordNotFound())

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

	}else{
		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()
	}
}