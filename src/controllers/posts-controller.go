package controllers

import "net/http"

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
	w.Write([]byte("Inserindo nova postagem!"))
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando postagem!"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo postagem!"))
}
