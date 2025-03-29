package menus

import (
	"encoding/json"
	"log"
	"net/http"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
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
		putting(w, r)
	} else if r.Method == "DELETE" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		deleting(w, r)
	} else {
		panic("method not allowed")
	}
}

func deleting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	validationDel(body)
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
	id1 := body["id"].(float64)
	id := uint64(id1)
	if err = repo.DelMenu(tx, id); err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationDel(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	if !arrayutils.Contains(keys, "id") {
		panic("bad: id is required")
	}
	id := body["id"].(float64)
	if id < 1 {
		panic("bad: id not found")
	}
}

func putting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	menu := repo.MapToMenu(body)
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
	log.Println("Menu", menu)
	err = repo.EditMenu(tx, menu, w)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func posting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	menu := repo.MapToMenu(body)
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
	log.Println("Menu", menu)
	err = repo.AddMenu(tx, menu, w)
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
