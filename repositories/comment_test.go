package repositories_test

import (
	"testing"

	"github.com/Lucky3028/try-go/models"
	"github.com/Lucky3028/try-go/repositories"
	_ "github.com/go-sql-driver/mysql"
)

func TestAddComment(t *testing.T) {
	comment := models.Comment{
		ArticleId: 1,
		Message:   "This comment is inserted by test",
	}
	newComment, err := repositories.AddComment(testDb, comment)
	if err != nil {
		t.Fatal(err)
	}

	expectedCommentId := 3
	if newComment.CommentId != expectedCommentId {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentId, newComment.CommentId)
	}

	t.Cleanup(func() {
		const query = `delete from comments where article_id = ? and message = ?;`
		testDb.Exec(query, comment.ArticleId, comment.Message)
	})
}

func TestListCommentsByArticleId(t *testing.T) {
	articleId := 1
	list, err := repositories.ListCommentsByArticleId(testDb, articleId)
	if err != nil {
		t.Fatal(err)
	}

	expectedLength := 2
	if actualLength := len(list); actualLength != expectedLength {
		t.Errorf("want %d but got %d comments\n", expectedLength, actualLength)
	}
	for _, comment := range list {
		if comment.ArticleId != articleId {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleId, comment.ArticleId)
		}
	}
}
