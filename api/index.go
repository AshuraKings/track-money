package handler

import (
	"errors"
	"net/http"
	"track/lib"
	"track/lib/repo"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		resp := map[string]interface{}{}
		resp["msg"] = "Hello World"
		roles, done := repo.AllRoles(w)
		if !done {
			return
		}
		resp["roles"] = roles
		lib.SendJson(resp, w)
	} else {
		panic(errors.New("method not allowed"))
	}
}
