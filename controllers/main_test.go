package controllers_test

import (
	"testing"

	"github.com/Lucky3028/try-go/controllers"
	"github.com/Lucky3028/try-go/controllers/testdata"
)

var articleController *controllers.ArticleController

func TestMain(m *testing.M) {
	service := testdata.NewServiceMock()
	articleController = controllers.NewArticleController(service)

	m.Run()
}
