package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/KTakao01/GoApi/models"
	"github.com/gorilla/mux"
)

// ハンドラの定義
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello,world!\n")
}

// POST /article　のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	//1.バイトスライスreqBodybufferをなんらかの形で用意

	var reqBodybuffer []byte

	//2.Readメソッドでリクエストボディを読み出し
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return

	}

	//3.ボディをCloseする
	defer req.Body.Close()

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// GET /article/{id}のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]
	articleID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid query parameter: %v", idParam), http.StatusBadRequest)
		return
	}
	//resString := fmt.Sprintf("Article No.%d\n", articleID)
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)

}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
