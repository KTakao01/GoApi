package services

import (
	"log"

	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	//記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	//コメント一覧を取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	//コメント一覧を記事詳細（Article構造体)に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostArticleHandlerで使うことを想定したサービス
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	//記事データをDBに挿入する
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	//挿入した記事を返す(取得する必要はない)
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	//指定ページの記事一覧をDBから取得する
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		log.Printf("Error selecting article list: %v", err)
		return nil, err
	}

	//取得した値を返す
	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	//指定IDの記事のいいね数を+1して結果を返却
	//いいね数を+1
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
