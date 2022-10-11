package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

const serverAddress = "http://localhost:8080"

func genUnexpectedStatusCodeMsg(expected int, actual int) string {
	return fmt.Sprintf("unexpected StatusCode: want %d but %d\n", expected, actual)
}

func TestListArticlesHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "alphabet", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/article/list?page=%s", serverAddress, tt.query)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			response := httptest.NewRecorder()
			articleController.ListArticlesHandler(response, request)

			if response.Code != tt.resultCode {
				t.Errorf(genUnexpectedStatusCodeMsg(tt.resultCode, response.Code))
			}
		})
	}
}

func TestArticleDetailHandler(t *testing.T) {
	var tests = []struct {
		name       string
		articleID  string
		resultCode int
	}{
		{name: "number pathparam", articleID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", articleID: "aaa", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("%s/article/%s", serverAddress, tt.articleID)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			response := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/article/{id:[1-9][0-9]*}", articleController.ArticleDetailHandler).Methods(http.MethodGet)
			router.ServeHTTP(response, request)

			if response.Code != tt.resultCode {
				t.Errorf(genUnexpectedStatusCodeMsg(tt.resultCode, response.Code))
			}
		})
	}
}
