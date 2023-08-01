package services

import (
	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/repositories"
)

// PostCommentHandlerで使用することを想定したサービス
func PostCommentService(comment models.Comment) (models.Comment, error) {

	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	//引数のコメント情報をDBに挿入
	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, err
}
