package models

import "time"

var (
	Comment1 = Comment{
		CommentId: 1,
		ArticleId: 1,
		Message:   "test comment 1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentId: 2,
		ArticleId: 1,
		Message:   "test comment 2",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		Id:          1,
		Title:       "article 1",
		Contents:    "This is the test 1st article.",
		UserName:    "Me",
		NiceCounts:  0,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
	Article2 = Article{
		Id:         2,
		Title:      "article 2",
		Contents:   "This is the test 2nd article.",
		UserName:   "Me",
		NiceCounts: 0,
		CreatedAt:  time.Now(),
	}
)
