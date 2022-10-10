package services

import (
	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
)

func GetArticle(id int) (models.Article, error) {
	db, err := connectToDb()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.FindArticleById(db, id)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.ListCommentsByArticleId(db, id)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = comments

	return article, nil
}

func PostArticle(article models.Article) (models.Article, error) {
	db, err := connectToDb()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.AddArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticlesList(page int) ([]models.Article, error) {
	db, err := connectToDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	list, err := repositories.ListArticles(db, page)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func IncrementNiceCounts(article models.Article) (models.Article, error) {
	db, err := connectToDb()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	if err := repositories.IncrementNiceCounts(db, article.Id); err != nil {
		return models.Article{}, err
	}

	result := article
	result.NiceCounts = article.NiceCounts + 1

	return result, nil
}
