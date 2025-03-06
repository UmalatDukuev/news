package repository

import (
	"fmt"

	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) Create(post models.Post) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, author_id) values ($1, $2, $3) RETURNING id", postsTable)

	row := r.db.QueryRow(query, post.Title, post.Description, post.AuthorID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostPostgres) Update(post models.Post) error {

	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id = $3", postsTable)
	row, err := r.db.Exec(query, post.Title, post.Description, post.ID)
	if err != nil {
		return err
	}
	fmt.Println(row.RowsAffected())

	return nil
}

func (r *PostPostgres) GetById(postID int) (models.Post, error) {
	var post models.Post
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postsTable)
	err := r.db.Get(&post, query, postID)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r *PostPostgres) GetAll() ([]models.Post, error) {
	var posts []models.Post
	query := fmt.Sprintf("SELECT * FROM %s", postsTable)
	err := r.db.Select(&posts, query)
	if err != nil {
		return []models.Post{}, err
	}
	return posts, nil
}
