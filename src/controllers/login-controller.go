package controllers

import (
	"api/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

	token, err := authentication.MakeToken(saveUser.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userId := strconv.FormatUint(saveUser.ID, 10)

	responses.JSON(w, http.StatusOK, models.DataAuth{Id: userId, Token: token})
}
