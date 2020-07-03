package main

import "net/http"

func newClient() *http.Client {
	return &http.Client{}
}
