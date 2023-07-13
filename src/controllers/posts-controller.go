package controllers

import (
	"api/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SelectPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	//fmt.Printf("SelecPosts %d", userId)

	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, err := repository.SelectAll(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
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
	//
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
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
	posts, err := repository.Search(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func InsertPost(w http.ResponseWriter, r *http.Request) {
	//
	userId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	//fmt.Printf("InsertPost %d\n", userId)

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
	//fmt.Printf("InsertPost %s\n", postName)

	repository := repositories.NewPostRepository(db)

	result, err := repository.CheckTitle(userId, post.Title)
	if err != nil || result {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	postId, err := repository.Insert(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	post.Id = postId

	responses.JSON(w, http.StatusCreated, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	//
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postSaved, err := repository.Select(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postSaved.AuthorId != userId {
		responses.Error(w, http.StatusForbidden,
			errors.New("você não pode atualizar uma postagem que não é sua"))
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

	if err = repository.Update(postId, post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	//
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userId, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.DBConnect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postSaved, err := repository.Select(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postSaved.AuthorId != userId {
		responses.Error(w, http.StatusForbidden,
			errors.New("você não pode excluir uma postagem que não é sua"))
		return
	}

	if err = repository.Delete(postId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
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
	if err = repository.Like(postId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
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
	if err = repository.Unlike(postId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
