package handler

import (
	"net/http"
	"track/lib"
	"track/lib/session"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{}) {
			return
		}
		lib.SendJson(map[string]any{"msg": "Success"}, w)
	} else {
		panic("method not allowed")
	}
}
