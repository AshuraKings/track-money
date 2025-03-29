package menus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	roleId1, menus1 := body["roleId"].(float64), body["menus"].([]any)
	roleId, menus := uint64(roleId1), arrayutils.Map(menus1, func(v any, _ int) uint64 { return uint64(v.(float64)) })
	insertQuery(w, tx, "DELETE FROM role_has_menu WHERE role_id=$1", roleId)
	for _, m := range menus {
		insertQuery(w, tx, "INSERT INTO role_has_menu(role_id,menu_id) VALUES($1,$2)", roleId, m)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func insertQuery(w http.ResponseWriter, tx *sql.Tx, query string, args ...any) {
	stmt, err := tx.Prepare(query)
	defer lib.StmtClose(w, stmt)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
}

func validationPost(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	for _, k := range []string{"roleId", "menus"} {
		if !arrayutils.Contains(keys, k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	roleId, menus := body["roleId"].(float64), body["menus"].([]any)
	if roleId < 1 {
		panic("bad: roleId not found")
	}
	if len(menus) < 1 {
		panic("bad: menus is empty")
	}
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
