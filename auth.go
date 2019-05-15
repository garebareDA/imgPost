package main

import(
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"log"
	"net/http"
	"context"
)

func auth(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func callBack(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Println(c.Writer, err)
		return
	}

	log.Printf("%#v", user)
}

func logOut(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func contextWithProviderName(c *gin.Context, provider string) (*http.Request){
	return  c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}