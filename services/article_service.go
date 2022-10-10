package services

import (
	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
)

func (service *ApplicationService) GetArticle(id int) (models.Article, error) {
	article, err := repositories.FindArticleById(service.db, id)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.ListCommentsByArticleId(service.db, id)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = comments

	return article, nil
}

func (service *ApplicationService) PostArticle(article models.Article) (models.Article, error) {
	newArticle, err := repositories.AddArticle(service.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func (service *ApplicationService) GetArticlesList(page int) ([]models.Article, error) {
	list, err := repositories.ListArticles(service.db, page)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *ApplicationService) IncrementNiceCounts(article models.Article) (models.Article, error) {
	if err := repositories.IncrementNiceCounts(service.db, article.Id); err != nil {
		return models.Article{}, err
	}

	result := article
	result.NiceCounts = article.NiceCounts + 1

	return result, nil
}
