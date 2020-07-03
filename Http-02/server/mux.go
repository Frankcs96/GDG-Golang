package main

import "net/http"

func newMux() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", Cached(loginHandler))
	mux.HandleFunc("/cache", cacheHandler)

	return mux
}
