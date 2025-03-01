package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"

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
		login(w, body)
	} else {
		panic(errors.New("method not allowed"))
	}
}

func login(w http.ResponseWriter, body map[string]any) {
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
	username := body["username"].(string)
	user, err := repo.UserByUsername(tx, username)
	if err != nil {
		panic(err)
	}
	password := body["password"].(string)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic(err)
	}
	session.PutSessionToResponse(w, user.Id)
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validation(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	if !arrayutils.Contains(keys, "username") {
		panic("username is required")
	}
	if !arrayutils.Contains(keys, "password") {
		panic("password is required")
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
