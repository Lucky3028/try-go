package main

import (
	"log"
	"net/http"

	"github.com/Lucky3028/try-go/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", controllers.HelloHandler).Methods(http.MethodGet)
	router.HandleFunc("/article", controllers.PostArticleHandler).Methods(http.MethodPost)
	router.HandleFunc("/article/list", controllers.ListArticlesHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/{id:[1-9][0-9]*}", controllers.ArticleDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/nice", controllers.PostNiceHandler).Methods(http.MethodPost)
	router.HandleFunc("/comment", controllers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
