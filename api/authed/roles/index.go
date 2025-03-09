package handler

import (
	"net/http"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"

	arrayutils "github.com/AchmadRifai/array-utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		getting(w)
	} else if r.Method == "POST" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		lib.SendJson(map[string]any{}, w)
	} else if r.Method == "PUT" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		lib.SendJson(map[string]any{}, w)
	} else if r.Method == "DELETE" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		lib.SendJson(map[string]any{}, w)
	} else {
		panic("method not allowed")
	}
}

func getting(w http.ResponseWriter) {
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
	lib.SendJson(map[string]any{
		"msg": "Success",
		"roles": arrayutils.Map(roles, func(v repo.Role, _ int) map[string]any {
			return map[string]any{
				"id":   v.Id,
				"name": v.Nm,
			}
		}),
	}, w)
}
