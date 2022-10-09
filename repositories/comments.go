package repositories

import (
	"database/sql"

	"github.com/Lucky3028/try-go/models"
)

func AddComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const query = `
		insert into comments (article_id, comments, created_at) values (?, ?, now());
	`
	result, err := db.Exec(query)
	if err != nil {
		return models.Comment{}, nil
	}

	id, _ := result.LastInsertId()
	newComment := models.Comment{
		CommentId: int(id),
		ArticleId: comment.ArticleId,
		Message:   comment.Message,
	}

	return newComment, nil
}

func ListCommentsByArticleId(db *sql.DB, articleId int) ([]models.Comment, error) {
	const query = `select comment_id, message, created_at from comments where article_id = ?;`
	rows, err := db.Query(query, articleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentList := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		comment.ArticleId = articleId
		rows.Scan(&comment.CommentId, &comment.Message, &comment.CreatedAt)
		commentList = append(commentList, comment)
	}

	return commentList, nil
}
