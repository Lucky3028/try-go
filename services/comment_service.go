package services

import (
	"github.com/Lucky3028/try-go/app_errors"
	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
)

func (service *ApplicationService) PostComment(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.AddComment(service.db, comment)
	if err != nil {
		err = app_errors.InsertDataFailed.Wrap(err, "fail to record data")

		return models.Comment{}, err
	}

	return newComment, nil
}
