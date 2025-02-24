package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"track/lib"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "POST" {
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			panic(err)
		}
		validation(body)
		resp := map[string]interface{}{}
		resp["msg"] = "Success"
		lib.SendJson(resp, w)
	} else {
		panic(errors.New("method not allowed"))
	}
}

func validation(body map[string]interface{}) {
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
