package services

import (
	"database/sql"
	"errors"

	"github.com/Lucky3028/try-go/app_errors"
	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
)

func (service *ApplicationService) GetArticle(id int) (models.Article, error) {
	article, err := repositories.FindArticleById(service.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = app_errors.DataNotFound.Wrap(err, "no data")

			return models.Article{}, err
		}

		err = app_errors.GetDataFailed.Wrap(err, "fail to get data")

		return models.Article{}, err
	}

	comments, err := repositories.ListCommentsByArticleId(service.db, id)
	if err != nil {
		err = app_errors.GetDataFailed.Wrap(err, "fail to get data")

		return models.Article{}, err
	}

	article.CommentList = comments

	return article, nil
}

func (service *ApplicationService) PostArticle(article models.Article) (models.Article, error) {
	newArticle, err := repositories.AddArticle(service.db, article)
	if err != nil {
		err = app_errors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (service *ApplicationService) GetArticlesList(page int) ([]models.Article, error) {
	list, err := repositories.ListArticles(service.db, page)
	if err != nil {
		err = app_errors.GetDataFailed.Wrap(err, "fail to get data")

		return nil, err
	}
	if len(list) == 0 {
		err = app_errors.DataNotFound.Wrap(ErrDataNotFound, "no data")

		return nil, err
	}

	return list, nil
}

func (service *ApplicationService) IncrementNiceCounts(article models.Article) (models.Article, error) {
	if err := repositories.IncrementNiceCounts(service.db, article.Id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = app_errors.DataNotFound.Wrap(err, "no data")

			return models.Article{}, err
		}

		err = app_errors.UpdateDataFailed.Wrap(err, "fail to increment nice counts")

		return models.Article{}, err
	}

	result := article
	result.NiceCounts = article.NiceCounts + 1

	return result, nil
}
