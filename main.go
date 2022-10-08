package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"try-go/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	router.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	router.HandleFunc("/article/list", handlers.ListArticlesHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/{id:[1-9][0-9]*}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	router.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
