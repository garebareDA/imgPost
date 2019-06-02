package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"net/http"
	"log"
)

func Home(c *gin.Context) {
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