package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ハンドラの定義
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello,world!\n")
}
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

// GET /article/{id}のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]
	articleID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid query parameter: %v", idParam), http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
