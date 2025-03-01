package handler

import (
	"errors"
	"net/http"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		resp := map[string]any{}
		resp["msg"] = "Hello World"
		db, err := db.DbConn()
		defer lib.CloseDb(w, db)
		if err != nil {
			panic(err)
		}
		tx, err := db.Begin()
		defer lib.TxClose(tx, w)
		if err != nil {
			panic(err)
		}
		roles, err := repo.AllRoles(tx)
		if err != nil {
			panic(err)
		}
		resp["roles"] = roles
		lib.SendJson(resp, w)
	} else {
		panic(errors.New("method not allowed"))
	}
}
