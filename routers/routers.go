package routers

import (
	"net/http"

	"github.com/Lucky3028/try-go/controllers"
	"github.com/gorilla/mux"
)

func NewRouter(controller *controllers.ApplicationController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hello", controller.HelloHandler).Methods(http.MethodGet)

	router.HandleFunc("/article", controller.PostArticleHandler).Methods(http.MethodPost)
	router.HandleFunc("/article/list", controller.ListArticlesHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/{id:[1-9][0-9]*}", controller.ArticleDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/nice", controller.PostNiceHandler).Methods(http.MethodPost)

	router.HandleFunc("/comment", controller.PostCommentHandler).Methods(http.MethodPost)

	return router
}
