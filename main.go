package main

import (
	"log"
	"net/http"
	"try-go/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ListArticlesHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
