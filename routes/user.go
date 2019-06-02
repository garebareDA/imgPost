package routes

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func User(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", gin.H{
	})
}