package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"net/http"
	"imgPost/database"
	"log"
)

func User(c *gin.Context) {
	db := database.ConnectDB()
	defer db.Close()

	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	token := csrf.GetToken(c)

	userPage := database.UserData{}

	db.Where("user_id = ?", c.Param("user")).First(&userPage)

	log.Println(userID)
	log.Println(alive)
	c.HTML(http.StatusOK, "user.html", gin.H{
		"alive": alive,
		"id":userID,
		"name":userPage.UserName,
		"_csrf": token,
		"userID": userPage.UserID,
	})
}