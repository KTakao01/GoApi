package services

import (
	"database/sql"
	"errors"

	"github.com/KTakao01/GoApi/apperrors"
	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}

	articleChan := make(chan articleResult)
	defer close(articleChan)

	//記事の詳細を取得
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}

	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		//コメント一覧を取得
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err

		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}

	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
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
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	//挿入した記事を返す(取得する必要はない)
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	//指定ページの記事一覧をDBから取得する
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
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
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target data")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
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
