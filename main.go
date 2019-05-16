package main

import(
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/markbates/goth/providers/google"
	"imgPost/auth"
)

func main() {

	goth.UseProviders(
		google.New("token", "secret", "http://localhost:8000/auth/google/callback"),
	)

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("postSession", store))
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", home)
	router.GET("/auth/:provider", auth.Auth)
	router.GET("/auth/:provider/callback", auth.CallBack)
	router.GET("/auth/:provider/logout", auth.LogOut)

	router.POST("/upload", imagePost)

	router.Run(":8000")
}