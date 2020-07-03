package main

import (
	"encoding/json"
	"net/http"

	"github.com/fcsuarez96/http/user"
)

var (
	cache = make(map[string]string)
)

func Cached(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		var user user.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if _, ok := cache[user.Username]; ok {

		} else {

			db := NewDatabase()
			exist := db.CheckUser(user)
			if exist {
				cache[user.Username] = user.Password
			}

		}

	}
}
