package handler

import (
	"encoding/json"
	"fmt"
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
		lib.SendJson(map[string]any{}, w)
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
	id, name := body["id"].(float64), body["name"].(string)
	role, err := repo.RoleById(tx, uint64(id))
	if err != nil {
		panic(err)
	}
	log.Print("Role", role)
	role.Nm = name
	err = repo.EditRole(tx, role)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationPut(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	for _, v := range []string{"id", "name"} {
		if !arrayutils.Contains(keys, v) {
			panic(fmt.Sprintf("bad: %s is required", v))
		}
	}
	id, name := body["id"].(float64), body["name"].(string)
	if id < 1 {
		panic("bad: id not found")
	}
	if name == "" {
		panic("bad: name is required")
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
	err = repo.AddRole(tx, body["name"].(string))
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationPost(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	if !arrayutils.Contains(keys, "name") {
		panic("bad: name is required")
	}
	name := body["name"].(string)
	if name == "" {
		panic("bad: name is required")
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
