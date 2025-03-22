package handler

import (
	"encoding/json"
	"fmt"
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
		if !session.ValidationRole(w, r, []string{"admin", "fin"}) {
			return
		}
		getting(w)
	} else if r.Method == "POST" {
		if !session.ValidationRole(w, r, []string{"admin", "fin"}) {
			return
		}
		posting(w, r)
	} else if r.Method == "PUT" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		putting(w, r)
	} else {
		panic("method not allowed")
	}
}

func putting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	validationPut(body)
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
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationPut(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	for _, k := range []string{"id", "nm", "balance"} {
		if !arrayutils.Contains(keys, k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	nm, balance, id := body["nm"].(string), body["balance"].(float64), body["id"].(float64)
	if nm == "" {
		panic("bad: nm is required")
	}
	if balance < 0 {
		panic("bad: balance must be positive or zero")
	}
	if id < 1 {
		panic("bad: id not found")
	}
}

func posting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	validationPost(body)
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
	err = repo.AddWallet(tx, repo.FromMapToWallet(body))
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationPost(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	for _, k := range []string{"nm", "balance"} {
		if !arrayutils.Contains(keys, k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	nm, balance := body["nm"].(string), body["balance"].(float64)
	if nm == "" {
		panic("bad: nm is required")
	}
	if balance < 0 {
		panic("bad: balance must be positive or zero")
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
	wallets, err := repo.AllWallet(tx)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success", "wallets": wallets}, w)
}
