package handler

import (
	"net/http"
	"track/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		lib.SendJson(map[string]any{}, w)
	} else {
		panic("method not allowed")
	}
}
