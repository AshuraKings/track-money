package handler

import (
	"encoding/json"
	"log"
	"net/http"
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
		getting(w)
	} else if r.Method == "POST" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		posting(w, r)
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

func posting(w http.ResponseWriter, r *http.Request) {
	var body repo.Menu
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
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
	log.Println("Menu", body)
	err = repo.AddMenu(tx, body)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
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
	menus, err := repo.AllMenus(tx)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success", "menus": menus}, w)
}
