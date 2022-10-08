package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Hello World!\n")
}

func PostArticleHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Posting Article...\n")
}

func ListArticlesHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Article List\n")
}

func ArticleDetailHandler(writer http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(writer, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	body := fmt.Sprintf("Article No.%d\n", articleId)
	io.WriteString(writer, body)
}

func PostNiceHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Posting Nice...\n")
}

func PostCommentHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Posting Comment...\n")
}
