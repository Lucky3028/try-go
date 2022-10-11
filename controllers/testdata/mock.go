package testdata

import "github.com/Lucky3028/try-go/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (service *serviceMock) PostArticle(article models.Article) (models.Article, error) {
	return articlesTestData[1], nil
}

func (service *serviceMock) GetArticlesList(page int) ([]models.Article, error) {
	return articlesTestData, nil
}

func (service *serviceMock) GetArticle(articleId int) (models.Article, error) {
	return articlesTestData[0], nil
}

func (service *serviceMock) IncrementNiceCounts(article models.Article) (models.Article, error) {
	return articlesTestData[0], nil
}

func (service *serviceMock) PostComment(comment models.Comment) (models.Comment, error) {
	return commentsTestData[0], nil
}
