package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
	"golang.org/x/crypto/bcrypt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "POST" {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			panic(err)
		}
		validation(body)
		registering(w, body)
	} else {
		panic(errors.New("method not allowed"))
	}
}

func registering(w http.ResponseWriter, body map[string]any) {
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
	password := body["password"].(string)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		panic(err)
	}
	user := repo.NewUser(body["nm"].(string), body["username"].(string), string(bytes), 4)
	id, err := repo.AddUser(tx, user)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success", "id": id}, w)
}

func validation(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	if !arrayutils.Contains(keys, "nm") {
		panic("nm is required")
	}
	if !arrayutils.Contains(keys, "username") {
		panic("username is required")
	}
	if !arrayutils.Contains(keys, "password") {
		panic("password is required")
	}
	nm := body["nm"].(string)
	if nm == "" {
		panic("nm is required")
	}
	if len(nm) < 5 {
		panic("nm length is min 5")
	}
	username := body["username"].(string)
	if username == "" {
		panic("username is required")
	}
	if len(username) < 5 {
		panic("username length is min 5")
	}
	password := body["password"].(string)
	if password == "" {
		panic("password is required")
	}
	if len(password) < 8 {
		panic("password length is min 8")
	}
}
