package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Lucky3028/try-go/models"
	"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Hello World!\n")
}

func PostArticleHandler(writer http.ResponseWriter, req *http.Request) {
	var requestedArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&requestedArticle); err != nil {
		http.Error(writer, "fail to decode json\n", http.StatusInternalServerError)
		return
	}

	article := requestedArticle
	json.NewEncoder(writer).Encode(article)
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
	// TODO: use page
	log.Println(page)

	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(writer).Encode(&articles)
}

func ArticleDetailHandler(writer http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(writer, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	// TODO: use articleId
	log.Println(articleId)

	article := models.Article1
	json.NewEncoder(writer).Encode(&article)
}

func PostNiceHandler(writer http.ResponseWriter, req *http.Request) {
	article := models.Article1
	json.NewEncoder(writer).Encode(&article)
}

func PostCommentHandler(writer http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	json.NewEncoder(writer).Encode(&comment)
}
