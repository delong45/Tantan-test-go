package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tantan/user"
	"tantan/utils"

	"github.com/gorilla/mux"
)

type UserResp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func UserList(w http.ResponseWriter, r *http.Request) {
	var users []user.User
	users, err := user.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	var us []UserResp
	for _, u := range users {
		ur := UserResp{
			Id:   utils.GetString(u.Id),
			Name: u.Name,
			Type: u.Type,
		}
		us = append(us, ur)
	}

	response, err := json.Marshal(us)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "%s", response)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	var u user.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	err = u.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	ur := UserResp{
		Id:   utils.GetString(u.Id),
		Name: u.Name,
		Type: u.Type,
	}

	response, err := json.Marshal(ur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "%s", response)
}

func UserProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		UserList(w, r)
	} else if r.Method == "POST" {
		UserCreate(w, r)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Users do not support such method\n")
	}
}

func RelationList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Relationships do not support such method\n")
		return
	}

	vars := mux.Vars(r)
	user_id := vars["user_id"]
	rs, err := user.GetRelations(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	response, err := json.Marshal(rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "%s", response)
}

func RelationCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Relation Creator do not support such method\n")
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	var s user.State
	err = json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	vars := mux.Vars(r)
	user_id := vars["user_id"]
	other_user_id := vars["other_user_id"]

	relation, err := user.GreateRelation(user_id, other_user_id, s.State)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}

	response, err := json.Marshal(relation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Server internal error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "%s", response)
}
