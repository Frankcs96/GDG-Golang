package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fcsuarez96/http/user"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user user.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if user.Username == "" || user.Password == "" {
		fmt.Fprintf(w, "Empty fields not allowed\n")
		return
	}

	db := NewDatabase()
	userExist := db.CheckUser(user)

	if userExist {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}

}

func cacheHandler(w http.ResponseWriter, r *http.Request) {
	var user user.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if _, ok := cache[user.Username]; ok {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)

	}

}
