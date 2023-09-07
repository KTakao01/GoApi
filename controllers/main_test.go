package controllers_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/KTakao01/GoApi/controllers"
	"github.com/KTakao01/GoApi/services"
	_ "github.com/go-sql-driver//mysql"
)

// 1.テストに使うリソース（コントローラ構造体)を用意
var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	dbUSer := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	ser := services.NewMyAppService(db)
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
