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
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(writer, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	body := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(writer, body)
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
