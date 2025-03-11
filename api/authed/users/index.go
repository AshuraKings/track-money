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
	"golang.org/x/crypto/bcrypt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		getting(w)
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
	password := body["password"].(string)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		panic(err)
	}
	idRole := body["role"].(float64)
	role, err := repo.RoleById(tx, uint64(idRole))
	if err != nil {
		panic(err)
	}
	log.Println("Role", role)
	_, err = repo.AddUser(tx, repo.NewUser(body["name"].(string), body["username"].(string), string(bytes), uint64(idRole)))
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationPost(body map[string]any) {
	log.Println("Req", body)
	keys := mapsutils.KeysOfMap(body)
	for _, k := range []string{"name", "username", "password", "role"} {
		if !arrayutils.Contains(keys, k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	name, username, password, role := body["name"].(string), body["username"].(string), body["password"].(string), body["role"].(float64)
	if name == "" {
		panic("bad: name is required")
	}
	if len(username) < 5 {
		panic("bad: name min 5")
	}
	if len(password) < 8 {
		panic("bad: name min 8")
	}
	if role < 1 {
		panic("bad: role not found")
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
	users, err := repo.AllUsers(tx)
	if err != nil {
		panic(err)
	}
	roleMap := mapsutils.Map(arrayutils.Grouping(roles, func(v repo.Role, _ int) uint64 { return v.Id }), func(v []repo.Role, _ uint64) string {
		if len(v) > 0 {
			return v[0].Nm
		} else {
			return ""
		}
	})
	lib.SendJson(map[string]any{
		"msg": "Success",
		"users": arrayutils.Map(users, func(v repo.User, _ int) map[string]any {
			return map[string]any{
				"id":       v.Id,
				"name":     v.Nm,
				"username": v.Username,
				"role":     roleMap[v.RoleId],
			}
		}),
	}, w)
}
