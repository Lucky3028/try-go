package repositories_test

import (
	"testing"

	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
	"github.com/Lucky3028/try-go/repositories/testdata"
	_ "github.com/go-sql-driver/mysql"
)

func TestFindArticleById(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticlesTestData[0],
		},
		{
			testTitle: "subtest2",
			expected:  testdata.ArticlesTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			actual, err := repositories.FindArticleById(testDb, test.expected.Id)
			if err != nil {
				t.Fatal(err)
			}

			if actual.Id != test.expected.Id {
				t.Errorf("ID: get %d but want %d\n", actual.Id, test.expected.Id)
			}
			if actual.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", actual.Title, test.expected.Title)
			}
			if actual.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", actual.Contents, test.expected.Contents)
			}
			if actual.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", actual.UserName, test.expected.UserName)
			}
			if actual.NiceCounts != test.expected.NiceCounts {
				t.Errorf("NiceCounts: get %d but want %d\n", actual.NiceCounts, test.expected.NiceCounts)
			}
		})
	}
}

func TestListArticles(t *testing.T) {
	expected := len(testdata.ArticlesTestData)
	list, err := repositories.ListArticles(testDb, 1)
	if err != nil {
		t.Fatal(err)
	}

	if actual := len(list); actual != expected {
		t.Errorf("want %d but got %d articles\n", expected, actual)
	}
}

func TestAddArticle(t *testing.T) {
	article := models.Article{
		Title:    "Inserted by test",
		Contents: "This post is inserted by test",
		UserName: "go test",
	}
	newArticle, err := repositories.AddArticle(testDb, article)
	if err != nil {
		t.Fatal(err)
	}

	expectedArticleId := len(testdata.ArticlesTestData) + 1
	if newArticle.Id != expectedArticleId {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleId, newArticle.Id)
	}

	t.Cleanup(func() {
		const query = `delete from articles where title = ? and contents = ? and username = ?;`
		testDb.Exec(query, article.Title, article.Contents, article.UserName)
	})
}

func TestIncrementNiceCount(t *testing.T) {
	articleId := testdata.ArticlesTestData[0].Id
	before, err := repositories.FindArticleById(testDb, articleId)
	if err != nil {
		t.Fatal(err)
	}
	if err := repositories.IncrementNiceCounts(testDb, articleId); err != nil {
		t.Fatal(err)
	}
	after, err := repositories.FindArticleById(testDb, articleId)
	if err != nil {
		t.Fatal(err)
	}

	if after.NiceCounts-before.NiceCounts != 1 {
		t.Error("fail to update nice counts")
	}
}
