package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repository posts) Select(ID uint64) (models.Post, error) {
	//
	rows, err := repository.db.Query(
		`SELECT * FROM Posts WHERE Id = ?`, ID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {

		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorName,
			&post.Likes,
			&post.RegDate,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository posts) Insert(post models.Post) (uint64, error) {
	//
	statement, err := repository.db.Prepare(
		"INSERT INTO Posts (Title, Content, AuthorId, AuthorName, RegDate) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	regDate, err := time.Parse("02/01/2006", post.RegDate)
	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(post.Title, post.Content, post.AuthorId, post.AuthorName, regDate)
	if err != nil {
		return 0, err
	}

	fmt.Println(result.RowsAffected())
	return 1, nil
}

func (repository posts) CheckTitle(userId uint64, title string) (bool, error) {
	//
	row, err := repository.db.Query(
		"SELECT Id FROM Posts WHERE AuthorId = ? AND Title LIKE ?", userId, title,
	)
	if err != nil {
		return false, err
	}
	defer row.Close()

	var post models.Post
	if row.Next() {
		if err = row.Scan(&post.Id); err != nil {
			return false, err
		}
	}

	postId := &post.Id
	PID := *postId

	if PID > 0 {
		return true, errors.New("esse título já existe na postagem")
	}

	return false, nil
}
