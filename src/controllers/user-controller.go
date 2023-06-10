package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func SelectUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionando todos os usuários!"))
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionado um usuário!"))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	//
	criteria := strings.ToLower(r.URL.Query().Get("user"))
	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, err := repository.Search(criteria)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	//
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.CheckUser(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userID, err := repository.Insert(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	//w.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))

	user.ID = userID
	responses.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário!"))
}
