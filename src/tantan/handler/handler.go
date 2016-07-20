package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tantan/user"
)

func UserList(w http.ResponseWriter, r *http.Request) {
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

	response, err := json.Marshal(u)
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
}

func RelationCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Relation Create!\n"))
}
