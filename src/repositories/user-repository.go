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
