package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"log"
	"strconv"
)

var num int

//ImagePost アップロード画像の保存
func ImagePost(c *gin.Context) {
	imageupload.LimitFileSize(5242880, c.Writer, c.Request)

	img, err := imageupload.Process(c.Request, "file")
	if err != nil {
		log.Panicln(err)
	}

	if num == 0 {
		num = 1
	}else{
		num++
	}

	log.Println(num)
	img.Save("./img/" + strconv.Itoa(num) + ".jpg")
}