package repositories

import (
	"api/src/models"
	"database/sql"
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

func (repository posts) Insert(post models.Post) (uint64, error) {
	//
	statement, err := repository.db.Prepare(
		"INSERT INTO Posts (Title, Content, AuthorId, RegDate) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	regDate, err := time.Parse("02/01/2006", post.RegDate)
	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(post.Title, post.Content, post.AuthorId, regDate)
	if err != nil {
		return 0, err
	}

	fmt.Println(result.RowsAffected())
	return 1, nil
}
