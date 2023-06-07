package controllers

import "net/http"

func SelectUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionando todos os usuários!"))
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selecionado um usuário!"))
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Inserindo usuário!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário!"))
}
