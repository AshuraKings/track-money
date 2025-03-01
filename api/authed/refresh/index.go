package handler

import (
	"net/http"
	"track/lib"
	"track/lib/session"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		session.ValidationRole(w, r, []string{})
		lib.SendJson(map[string]any{"msg": "Success"}, w)
	} else {
		panic("method not allowed")
	}
}
