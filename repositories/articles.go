package repositories

import (
	"database/sql"

	"github.com/Lucky3028/try-go/models"
)

func AddArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const query = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`
	result, err := db.Exec(query, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()
	newArticle := article
	newArticle.Id = int(id)

	return newArticle, nil
}

const articlesPerPage = 5

func ListArticles(db *sql.DB, page int) ([]models.Article, error) {
	const query = `select * from articles limit ? offset ?;`
	rows, err := db.Query(query, articlesPerPage, ((page - 1) * articlesPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleList := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdAt sql.NullTime
		rows.Scan(&article.Id, &article.Title, &article.Contents, &article.UserName, &article.NiceCounts, &createdAt)
		if createdAt.Valid {
			article.CreatedAt = createdAt.Time
		}

		articleList = append(articleList, article)
	}

	return articleList, nil
}

func FindArticleById(db *sql.DB, id int) (models.Article, error) {
	const query = `select * from articles where article_id = ?;`
	row := db.QueryRow(query, id)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdAt sql.NullTime
	if err := row.Scan(&article.Id, &article.Title, &article.Contents, &article.UserName, &article.NiceCounts, &createdAt); err != nil {
		return models.Article{}, err
	}
	if createdAt.Valid {
		article.CreatedAt = createdAt.Time
	}

	return article, nil
}

func IncrementNiceCounts(db *sql.DB, articleId int) error {
	// NOTE: 本来は`update`1回だけで完結するが、トランザクションの練習のためかSQL文が2個に分けられているので、そのように実装する
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const selectQuery = `select nice from articles where article_id = ?;`
	row := tx.QueryRow(selectQuery, articleId)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	var oldNiceCounts int
	err = row.Scan(&oldNiceCounts)
	if err != nil {
		tx.Rollback()
		return err
	}

	newNiceCounts := oldNiceCounts + 1
	const updateQuery = `update articles set nice = ? where article_id = ?;`
	_, err = tx.Exec(updateQuery, newNiceCounts, articleId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
