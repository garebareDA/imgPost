package main

import(
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"sort"
)

func main() {

	goth.UseProviders(
		google.New("token", "secret", "http://localhost:8000/auth/google/callback"),
	)

	m := make(map[string]string)
	m["twitter"] = "twitter"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", home)
	router.GET("/auth/:provider", auth)
	router.GET("/auth/:provider/callback", callBack)
	router.GET("/auth/:provider/logout", logOut)

	router.Run(":8000")
}