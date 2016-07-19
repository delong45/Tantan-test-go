package handler

import "github.com/gorilla/mux"

func Init() *mux.Router {
	r := mux.NewRouter()

	s := r.Methods("GET", "POST", "PUT").Subrouter()
	s.HandleFunc("/users", UserProcess)
	s.HandleFunc("/users/{user_id:[0-9]+}/relationships", RelationList)
	s.HandleFunc("/users/{user_id:[0-9]+}/relationships/{other_user_id:[0-9]+}", RelationCreate)

	return s
}
