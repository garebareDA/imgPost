package main

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"log"
)

func home(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")
	log.Println(userID)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"h": "<h1>aaaaaaaaaaaaaaaaa</h1>",
		"alive": alive,
		"id":userID,
	})
}