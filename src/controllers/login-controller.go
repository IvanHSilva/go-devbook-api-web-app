package controllers

import (
	"api/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
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
	saveUser, err := repository.CheckMail(user.EMail)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPass(saveUser.Pass, user.Pass); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		w.Write([]byte("Login falhou! Verifique seu e-mail e senha!"))
		return
	}

	token, _ := authentication.MakeToken(saveUser.ID)
	fmt.Println(token)

	w.Write([]byte("Login efetuado com sucesso!\n"))
	w.Write([]byte("Token: " + token))

}
