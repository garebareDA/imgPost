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

	db := database.ConnectDB()
	defer db.Close()

	if alive == true {
		db.Where("user_id = ?", userID.(string)).First(&user)
			if db.Where("user_id = ?", userID.(string)).First(&user).RecordNotFound() && alive == true{
			c.Redirect(http.StatusFound, "/auth/google/logout")
		}
	}

	log.Println(user.UserName)

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

func error(c* gin.Context, Error string) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"error": Error,
	})
}