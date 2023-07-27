package repositories_test

import (
	"testing"

	"github.com/KTakao01/GoApi/models"
	"github.com/KTakao01/GoApi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %d but want %d", num, expectedNum)
	}

}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 2,
		Message:   "abcde",
	}
	expectedCommentNum := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new comment id is expected %d but got %d", expectedCommentNum, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments 
			where article_id = ? and message = ?
		`

		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)

		testDB.Exec("alter table comments auto_increment = 1")

	})

}
