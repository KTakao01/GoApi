package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first Article",
		Contents:    "This is the test article.",
		UserName:    "ktakao01",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
	Article2 = Article{
		ID:        2,
		Title:     "second Article",
		Contents:  "This is the test article.",
		UserName:  "ktakao01",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
