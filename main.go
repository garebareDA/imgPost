package main

import(
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/markbates/goth/providers/google"
	"github.com/utrack/gin-csrf"
	"imgPost/auth"
	"imgPost/routes"
	"imgPost/database"
)

func main() {

	db := database.ConnectDB()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE = InnoDB CHARSET=utf8mb4",).AutoMigrate(&database.ImgPostData{})
	db.Set("gorm:table_options", "ENGINE = InnoDB CHARSET=utf8mb4",).AutoMigrate(&database.UserData{})

	goth.UseProviders(
		google.New("token", "secret", "http://localhost:8000/auth/google/callback"),
	)

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))

	router.Static("/img","./img")
	router.Static("/icon", "./icon")

	router.Use(sessions.Sessions("postSession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "imgPoster",
		ErrorFunc: func(c *gin.Context){
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	router.LoadHTMLGlob("templates/*.html")

	//ホーム
	router.GET("/", routes.Home)

	//Auth認証
	router.GET("/auth/:provider", auth.Auth)
	router.GET("/auth/:provider/callback", auth.CallBack)
	router.GET("/auth/:provider/logout", auth.LogOut)

	router.GET("/acount", routes.Acount)

	router.POST("/", routes.ImagePost)
	router.POST("/acount", routes.CreateAcount)

	router.Run(":8000")
}