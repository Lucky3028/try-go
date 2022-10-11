package testdata

import "github.com/Lucky3028/try-go/models"

var articlesTestData = []models.Article{
	models.Article{
		Id:         1,
		Title:      "firstPost",
		Contents:   "This is my first blog.",
		UserName:   "Me",
		NiceCounts: 0,
	},
	models.Article{
		Id:         2,
		Title:      "secondPost",
		Contents:   "This is my second blog.",
		UserName:   "Me",
		NiceCounts: 0,
	},
}

var commentsTestData = []models.Comment{
	models.Comment{
		CommentId: 1,
		ArticleId: articlesTestData[0].Id,
		Message:   "1 get",
	},
	models.Comment{
		CommentId: 2,
		ArticleId: articlesTestData[0].Id,
		Message:   "Welcome to the Internet.",
	},
}
