package main

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"net/http"
)

func imagePost(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}