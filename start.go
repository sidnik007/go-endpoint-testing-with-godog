package main

import (
	. "net/http"
	"strings"
)

func startPageHandler(w ResponseWriter, r *Request) {
	requestMethod := r.Method
	if !(strings.EqualFold(requestMethod, "GET") || strings.EqualFold(requestMethod, "HEAD")) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(StatusMethodNotAllowed)
		w.Write([]byte("405 Method Not Allowed"))
	} else {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(StatusOK)
		_, err := w.Write([]byte("TODO"))
		if err != nil {
			panic(err)
		}
	}
}
