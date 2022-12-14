package models

import "time"

type Comment struct {
	CommentId int       `json:"comment_id"`
	ArticleId int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	Id          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceCounts  int       `json:"nice_counts"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}
