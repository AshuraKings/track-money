package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		getting(w, r)
	} else if r.Method == "POST" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		posting(w, r)
	} else {
		panic("method not allowed")
	}
}

func posting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func getting(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if !query.Has("id") {
		panic("bad: id is required")
	}
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
	id1, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		panic(err)
	}
	id := uint64(id1)
	activatedMenus, err := repo.MenuByRoleId2(tx, id)
	if err != nil {
		panic(err)
	}
	menus, err := repo.AllMenus(tx)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success", "activatedMenus": activatedMenus, "menus": menus}, w)
}
