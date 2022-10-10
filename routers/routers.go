package routers

import (
	"database/sql"
	"net/http"

	"github.com/Lucky3028/try-go/controllers"
	"github.com/Lucky3028/try-go/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	service := services.NewApplicationService(db)
	article := controllers.NewArticleController(service)
	comment := controllers.NewCommentController(service)

	router := mux.NewRouter()

	router.HandleFunc("/hello", article.HelloHandler).Methods(http.MethodGet)

	router.HandleFunc("/article", article.PostArticleHandler).Methods(http.MethodPost)
	router.HandleFunc("/article/list", article.ListArticlesHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/{id:[1-9][0-9]*}", article.ArticleDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/nice", article.PostNiceHandler).Methods(http.MethodPost)

	router.HandleFunc("/comment", comment.PostCommentHandler).Methods(http.MethodPost)

	return router
}
