package services

import (
	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	//引数のコメント情報をDBに挿入
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, err
}
