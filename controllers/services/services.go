package services

import "github.com/Lucky3028/try-go/models"

type IArticleService interface {
	PostArticle(article models.Article) (models.Article, error)
	GetArticlesList(page int) ([]models.Article, error)
	GetArticle(articleId int) (models.Article, error)
	IncrementNiceCounts(article models.Article) (models.Article, error)
}

type ICommentService interface {
	PostComment(comment models.Comment) (models.Comment, error)
}
