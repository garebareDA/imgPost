package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func home(c *gin.Context) {
	log.Println(c.Request)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"h": "<h1>aaaaaaaaaaaaaaaaa</h1>",
	})
}