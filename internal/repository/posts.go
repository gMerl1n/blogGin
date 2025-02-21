package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1n/blog/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	postsTable = "posts"
)

type RepositoryPost struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewRepositoryPost(db *pgxpool.Pool, logger *logrus.Logger) *RepositoryPost {
	return &RepositoryPost{
		db:     db,
		logger: logger,
	}
}

func (r *RepositoryPost) CreatePost(ctx context.Context, title, body string) (int, error) {

	var postID int

	query := fmt.Sprintf(
		`INSERT INTO %s (title, body)
	 	 VALUES ($1, $2)
	 	 RETURNING id`,
		postsTable,
	)

	if err := r.db.QueryRow(
		ctx,
		query,
		title,
		body,
	).Scan(&postID); err != nil {
		return 0, err
	}

	return postID, nil

}

func (r *RepositoryPost) GetPostByID(ctx context.Context, postID int) (*domain.Post, error) {

	var post domain.Post

	query := `SELECT p.id, u.name, p.title, p.body, p.updatet_at, p.created_at 
		 	  FROM posts AS p
		 	  JOIN users AS u ON p.user_id = u.id
		 	  WHERE id = $1`

	if err := r.db.QueryRow(
		ctx,
		query,
		postID,
	).Scan(&post.ID, &post.Author, &post.Title, &post.Body, &post.UpdatedAt, &post.CreatedAt); err != nil {
		return nil, err
	}

	return &post, nil
}
