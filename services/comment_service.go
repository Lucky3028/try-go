package services

import (
	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
)

func PostComment(comment models.Comment) (models.Comment, error) {
	db, err := connectToDb()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.AddComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
