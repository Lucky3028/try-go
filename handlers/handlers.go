package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Lucky3028/try-go/models"
	"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Hello World!\n")
}

func PostArticleHandler(writer http.ResponseWriter, req *http.Request) {
	len, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(writer, "Cannot get Content-Length", http.StatusBadRequest)
		return
	}
	reqBodyBuffer := make([]byte, len)

	if _, err := req.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(writer, "failed to get request body\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var requestedArticle models.Article
	if err := json.Unmarshal(reqBodyBuffer, &requestedArticle); err != nil {
		http.Error(writer, "fail to decode json\n", http.StatusInternalServerError)
		return
	}

	article := requestedArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(writer, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
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

	articles := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articles)
	if err != nil {
		errMessage := fmt.Sprintf("fail to encode json (page: %d)\n", page)
		http.Error(writer, errMessage, http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}

func ArticleDetailHandler(writer http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(writer, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMessage := fmt.Sprintf("fail to encode json (ArticleId: %d)\n", articleId)
		http.Error(writer, errMessage, http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}

func PostNiceHandler(writer http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(writer, "fail to encode json", http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}

func PostCommentHandler(writer http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(writer, "fail to encode json", http.StatusInternalServerError)
		return
	}

	writer.Write(jsonData)
}
