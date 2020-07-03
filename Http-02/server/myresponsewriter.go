package main

import (
	"bytes"
	"net/http"
)

// We can pass the data from the decorator to the handler function with this custom responsewriter
type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}
