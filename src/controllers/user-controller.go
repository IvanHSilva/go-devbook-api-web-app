package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SelectUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionando todos os usuários!"))
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionado um usuário!"))
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	//
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUserRepository(db)
	repository.Insert(user)

	//w.Write([]byte("Inserindo usuário!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário!"))
}
