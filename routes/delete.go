package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"strconv"
	"imgPost/database"
	"log"
)

func Delete(c *gin.Context){
	session := sessions.Default(c)
	userID := session.Get("userId")
	alive := session.Get("alive")

	c.Request.ParseForm()
	id := c.Request.Form["ID"]
	reqUserID := c.Request.Form["userID"]

	isAlive(alive, c)

	db := database.ConnectDB()
	defer db.Close()

	deletePost := database.ImgPostData{}
	d, err := strconv.Atoi(id[0])
	if err != nil{
		panic(err)
	}

	log.Println(reqUserID[0])

	if reqUserID[0] == userID {
		db.Where("post_id = ?", d).First(&deletePost)
		db.Delete(&deletePost)
		c.Redirect(http.StatusFound, "/user/" + reqUserID[0])
		c.Abort()
	}else{
		c.Redirect(http.StatusFound, "/auth/google/logout")
		c.Abort()
	}

	}