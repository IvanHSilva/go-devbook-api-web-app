package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Select(ID uint64) (models.User, error) {
	//
	rows, err := repository.db.Query(
		"SELECT Id, Name, EMail, RegDate FROM Users WHERE Id = ?", ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.EMail,
			&user.RegDate,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Insert(user models.User) (uint64, error) {
	//
	//today := time.Now()
	//regDate := today.Format("01/02/2006") // Pt-BR
	//fmt.Println(today.Format("02/01/2006")) // USA

	statement, err := repository.db.Prepare("INSERT INTO Users (Name, EMail, Password, RegDate) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	regDate, err := time.Parse("02/01/2006", user.RegDate)
	if err != nil {
		log.Fatal(err)
	}
	result, err := statement.Exec(user.Name, user.EMail, user.Pass, regDate)
	if err != nil {
		return 0, err
	}
	fmt.Println(result.RowsAffected())
	return 1, nil

	// lastID, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }

	// return uint64(lastID), nil
}

func (repository users) Update(ID uint64, user models.User) error {
	//
	statement, err := repository.db.Prepare(
		"UPDATE Users SET Name = ?, EMail = ?, RegDate = ? WHERE Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	regDate, err := time.Parse("02/01/2006", user.RegDate)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := statement.Exec(user.Name, user.EMail, regDate, ID); err != nil {
		return err
	}
	return nil
}

func (repository users) Delete(ID uint64) error {
	//
	statement, err := repository.db.Prepare(
		"DELETE FROM Users WHERE Id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repository users) Search(criteria string) ([]models.User, error) {
	//
	criteria = fmt.Sprintf("%%%s%%", criteria)

	rows, err := repository.db.Query(
		"SELECT Id, Name, EMail, RegDate FROM Users WHERE Name LIKE ?", criteria,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.EMail,
			&user.RegDate,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) CheckMail(email string) (models.User, error) {
	//
	row, err := repository.db.Query(
		"SELECT Id, EMail, Password FROM Users WHERE EMail = ?", email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.EMail, &user.Pass); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Follow(userId, followerId uint64) error {
	//
	statement, err := repository.db.Prepare(
		"INSERT INTO FollowersUsers (UserId, FollowerId) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}
	return nil
}

func (repository users) Unfollow(userId, followerId uint64) error {
	//
	statement, err := repository.db.Prepare(
		"DELETE FROM FollowersUsers WHERE UserId = ? AND FollowerId = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}
	return nil
}
