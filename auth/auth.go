package auth

import(
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/gin-contrib/sessions"
	"log"
	"net/http"
	"context"
)

func Auth(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func CallBack(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Println(c.Writer, err)
		return
	}

	session := sessions.Default(c)
	session.Set("alive", true)
	session.Set("userId", user.UserID)
	session.Save()
}

func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	log.Println("Session clear")
}

func contextWithProviderName(c *gin.Context, provider string) (*http.Request){
	return  c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}