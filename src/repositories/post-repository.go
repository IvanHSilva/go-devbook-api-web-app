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

func (repository posts) SelectAll(userId uint64) ([]models.Post, error) {
	//
	rows, err := repository.db.Query(
		"SELECT * FROM Posts WHERE AuthorId = ? ORDER BY RegDate DESC", userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorName,
			&post.Likes,
			&post.RegDate,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
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

func (repository posts) Search(userId uint64) ([]models.Post, error) {
	//
	rows, err := repository.db.Query(
		"SELECT * FROM Posts WHERE AuthorId = ? ORDER BY RegDate DESC", userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.AuthorName,
			&post.Likes,
			&post.RegDate,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repository posts) Update(postId uint64, post models.Post) error {
	//
	statement, err := repository.db.Prepare(
		"UPDATE Posts SET Title = ?, Content = ?, RegDate = ? WHERE Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	regDate, err := time.Parse("02/01/2006", post.RegDate)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := statement.Exec(post.Title, post.Content, regDate, postId); err != nil {
		return err
	}
	return nil
}

func (repository posts) Delete(ID uint64) error {
	//
	statement, err := repository.db.Prepare(
		"DELETE FROM Posts WHERE Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}
	return nil
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

func (repository posts) Like(postId uint64) error {
	//
	statement, err := repository.db.Prepare(
		"UPDATE Posts SET Likes = Likes + 1 WHERE Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}
	return nil
}

func (repository posts) Unlike(postId uint64) error {
	//
	statement, err := repository.db.Prepare(
		"UPDATE Posts SET Likes = Likes - 1 WHERE Likes > 0  AND Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}
	return nil
}
