package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/services"
	"github.com/gorilla/mux"
)

// コントローラ構造体
type MyAppController struct {
	service *services.MyAppService
}

func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

// ハンドラの定義
func (c *MyAppController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello,world!\n")
}

// POST /article　のハンドラ
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	//JSONデコーダー
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	//JSONエンコーダー
	json.NewEncoder(w).Encode(article)
}

// GET /article/list のハンドラ
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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
	log.Printf("Page number: %d\n", page)
	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(articleList)
	if err != nil {
		log.Printf("Error encoding JSON: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GET /article/{id}のハンドラ
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]
	articleID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid query parameter: %v", idParam), http.StatusBadRequest)
		return
	}
	//resString := fmt.Sprintf("Article No.%d\n", articleID)
	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)

}

func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	//リクエスト情報をデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	//リクエスト情報に基づいてnice+1(DBデータから+1ではないことに注意)
	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exex\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
