package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fcsuarez96/http/user"
)

var (
	cache = make(map[string]string)
)

func Cached(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user user.User

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		mrw := &MyResponseWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}

		// read json
		err = json.Unmarshal(body, &user)

		if err != nil {
			http.Error(w, err.Error(), 500)
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

		h.ServeHTTP(mrw, r)

	}

}
