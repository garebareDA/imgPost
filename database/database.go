package database

import(
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

const(

	//Dialect
	Dialect = "mysql"

	//DBUser ユーザー名
	DBUser = "user"

	//DBPass パスワード
	DBPass = "pass"

	//DBProtocol プロトコル
	DBProtocol = "tcp(127.0.0.1)"

	//DBName DB名
	DBName = "imgPost"

	//DBchar 文字コード
	DBchar = "charset=utf8mb4"
)

//ConnectDB DBにアクセス
func ConnectDB() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s?%s&parseTime=true"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName, DBchar,)

	db, err := gorm.Open(Dialect, connect)
	if err != nil {
		panic(err)
	}

	return db
}