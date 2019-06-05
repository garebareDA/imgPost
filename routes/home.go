package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"net/http"
	"imgPost/database"
	"log"
)

func Home(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	token := csrf.GetToken(c)

	user := database.UserData{}

	if alive == true {
		db := database.ConnectDB()
		defer db.Close()
		user.UserID = userID.(string)
		db.First(&user)
	}

	log.Println(userID)
	log.Println(alive)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alive": alive,
		"id":userID,
		"name":user.UserName,
		"_csrf": token,
	})
}

func isAlive(alive interface{}, c *gin.Context) {
	log.Println(alive)
	if alive == nil{

		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()

	}else if alive.(bool) == true{
		return

	}else{

		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()
	}
}