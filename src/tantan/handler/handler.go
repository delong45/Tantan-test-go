package handler

import "net/http"

func UserList() {
}

func UserCreate() {
}

func UserProcess(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Process!\n"))
}

func RelationList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Relation List!\n"))
}

func RelationCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Relation Create!\n"))
}
