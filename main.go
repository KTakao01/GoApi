package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/KTakao01/GoApi/handlers"
)

func main() {

	r := mux.NewRouter()
	//定義したハンドラの登録
	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	//サーバー起動
	log.Fatal(http.ListenAndServe(":8080", r))
}
