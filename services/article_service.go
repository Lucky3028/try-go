package services

import (
	"database/sql"
	"errors"

	"github.com/Lucky3028/try-go/app_errors"
	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
)

func (service *ApplicationService) GetArticle(id int) (models.Article, error) {
	var article models.Article
	var comments []models.Comment
	var getArticleErr, getCommentsErr error

	type articleResult struct {
		article models.Article
		err     error
	}
	articleChannel := make(chan articleResult)
	defer close(articleChannel)

	go func(ch chan<- articleResult, db *sql.DB, articleId int) {
		article, err := repositories.FindArticleById(db, articleId)
		ch <- articleResult{article, err}
	}(articleChannel, service.db, id)

	type commentsResult struct {
		comments *[]models.Comment
		err      error
	}
	commentsChannel := make(chan commentsResult)
	defer close(commentsChannel)

	go func(ch chan<- commentsResult, db *sql.DB, commentId int) {
		comments, err := repositories.ListCommentsByArticleId(db, commentId)
		ch <- commentsResult{&comments, err}
	}(commentsChannel, service.db, id)

	for i := 0; i < 2; i++ {
		select {
		case articleResult := <-articleChannel:
			article, getArticleErr = articleResult.article, articleResult.err
		case commentsResult := <-commentsChannel:
			comments, getCommentsErr = *commentsResult.comments, commentsResult.err
		}
	}

	if getArticleErr != nil {
		if errors.Is(getArticleErr, sql.ErrNoRows) {
			err := app_errors.DataNotFound.Wrap(getArticleErr, "no data")

			return models.Article{}, err
		}

		err := app_errors.GetDataFailed.Wrap(getArticleErr, "fail to get data")

		return models.Article{}, err
	}

	if getCommentsErr != nil {
		err := app_errors.GetDataFailed.Wrap(getCommentsErr, "fail to get data")

		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)

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
			err = app_errors.NoTargetData.Wrap(err, "no data")

			return models.Article{}, err
		}

		err = app_errors.UpdateDataFailed.Wrap(err, "fail to increment nice counts")

		return models.Article{}, err
	}

	result := article
	result.NiceCounts = article.NiceCounts + 1

	return result, nil
}
