package imagePost

import(
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"log"
	"strconv"
)

var num int

//ImagePost アップロード画像の保存
func ImagePost(c *gin.Context) {
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
	log.Println("./img/a" + string(num) + ".jpg")
	img.Save("./img/a" + strconv.Itoa(num) + ".jpg")
}