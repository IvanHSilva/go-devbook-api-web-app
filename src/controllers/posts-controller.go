package controllers

import (
	"api/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SelectPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionando todas as postagens!"))
}

func SelectPost(w http.ResponseWriter, r *http.Request) {
	//
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post, err := repository.Select(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

func SearchPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Procurando postagens!"))
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	//
	userId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = post.CheckPost(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	post.AuthorId = userId
	repoUser := repositories.NewUserRepository(db)
	postName, err := repoUser.GetUserName(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	post.AuthorName = postName

	repository := repositories.NewPostRepository(db)
	postId, err := repository.Insert(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	post.Id = postId

	responses.JSON(w, http.StatusCreated, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando postagem!"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo postagem!"))
}