package services

import (
	"log"

	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {

	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	//記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	//コメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	//コメント一覧を記事詳細（Article構造体)に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostArticleHandlerで使うことを想定したサービス
func PostArticleService(article models.Article) (models.Article, error) {
	//DBに接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	//記事データをDBに挿入する
	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	//挿入した記事を返す(取得する必要はない)
	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	//DBに接続
	db, err := connectDB()
	if err != nil {
		log.Printf("Error connecting to DB: %v", err)
		return nil, err
	}
	defer db.Close()

	//指定ページの記事一覧をDBから取得する
	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		log.Printf("Error selecting article list: %v", err)
		return nil, err
	}

	//取得した値を返す
	return articleList, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	//DBに接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	//指定IDの記事のいいね数を+1して結果を返却
	//いいね数を+1
	updatedArticle, err := repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return updatedArticle, nil
}
