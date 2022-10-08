package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(writer, "Hello World!\n")
	} else {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(writer, "Posting Article...\n")
	} else {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func ListArticlesHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(writer, "Article List\n")
	} else {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(writer http.ResponseWriter, req *http.Request) {
	articleId := 1
	body := fmt.Sprintf("Article No.%d\n", articleId)

	if req.Method == http.MethodGet {
		io.WriteString(writer, body)
	} else {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostNiceHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(writer, "Posting Nice...\n")
	} else {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostCommentHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(writer, "Posting Comment...\n")
	} else {
		http.Error(writer, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
