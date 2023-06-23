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
)

func SelectPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionando todas as postagens!"))
}

func SelectPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionando uma postagens!"))
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

	post.AuthorId = userId

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
