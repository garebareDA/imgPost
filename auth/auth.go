package auth

import(
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/gin-contrib/sessions"
	"log"
	"net/http"
	"context"
	"imgPost/database"
)

//Auth 認証にリダイレクト
func Auth(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

//CallBack UserIDをDBに追加
func CallBack(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Println(c.Writer, err)
		return
	}

	//sessionに追加
	session := sessions.Default(c)
	session.Set("alive", true)
	session.Set("userId", user.UserID)
	session.Save()

	db := database.ConnectDB()
	defer db.Close()

	//DBにIDを追加
	db.Save(database.UserData{ID:user.UserID})
	//TODOアカウント登録画面に遷移
}

//LogOut sessionを削除
func LogOut(c *gin.Context) {
	session := sessions.Default(c)

	//sessionを削除
	session.Clear()
	session.Save()
	log.Println("Session clear")
}

//contextWithProviderName サポート関数
func contextWithProviderName(c *gin.Context, provider string) (*http.Request){
	return  c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}