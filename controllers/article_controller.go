package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Lucky3028/try-go/controllers/services"
	"github.com/Lucky3028/try-go/models"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.IArticleService
}

func NewArticleController(service services.IArticleService) *ArticleController {
	return &ArticleController{service}
}

func (controller *ArticleController) HelloHandler(writer http.ResponseWriter, req *http.Request) {
	io.WriteString(writer, "Hello World!\n")
}

func (controller *ArticleController) PostArticleHandler(writer http.ResponseWriter, req *http.Request) {
	var requestedArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&requestedArticle); err != nil {
		http.Error(writer, "fail to decode json\n", http.StatusInternalServerError)
		return
	}

	article, err := controller.service.PostArticle(requestedArticle)
	if err != nil {
		http.Error(writer, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(article)
}

func (controller *ArticleController) ListArticlesHandler(writer http.ResponseWriter, req *http.Request) {
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

	articles, err := controller.service.GetArticlesList(page)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(&articles)
}

func (controller *ArticleController) ArticleDetailHandler(writer http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(writer, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := controller.service.GetArticle(articleId)
	if err != nil {
		http.Error(writer, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(&article)
}

func (controller *ArticleController) PostNiceHandler(writer http.ResponseWriter, req *http.Request) {
	var requestedArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&requestedArticle); err != nil {
		http.Error(writer, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := controller.service.IncrementNiceCounts(requestedArticle)
	if err != nil {
		http.Error(writer, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(&article)
}
