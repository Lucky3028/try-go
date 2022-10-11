package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Lucky3028/try-go/app_errors"
	"github.com/Lucky3028/try-go/controllers/services"
	"github.com/Lucky3028/try-go/models"
)

type CommentController struct {
	service services.ICommentService
}

func NewCommentController(service services.ICommentService) *CommentController {
	return &CommentController{service}
}

func (controller *CommentController) PostCommentHandler(writer http.ResponseWriter, req *http.Request) {
	var requestedComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&requestedComment); err != nil {
		err = app_errors.RequestBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(writer, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := controller.service.PostComment(requestedComment)
	if err != nil {
		http.Error(writer, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(&comment)
}
