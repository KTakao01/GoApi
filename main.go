package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KTakao01/GoApi/controllers"
	"github.com/KTakao01/GoApi/services"
	_ "github.com/go-sql-driver/mysql"

	"github.com/KTakao01/GoApi/routers"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func main() {

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)
	r := routers.NewRouter(con)
	log.Println("server start at port 8080")
	//サーバー起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
