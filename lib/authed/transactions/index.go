package transactions

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"track/lib"
	"track/lib/db"
	"track/lib/session"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin", "fin"}) {
			return
		}
		getting(w, r)
	} else if r.Method == "POST" {
		if !session.ValidationRole(w, r, []string{"admin", "fin"}) {
			return
		}
		lib.SendJson(map[string]any{"msg": "Success"}, w)
	} else if r.Method == "DELETE" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		lib.SendJson(map[string]any{"msg": "Success"}, w)
	} else {
		panic("method not allowed")
	}
}

func getting(w http.ResponseWriter, r *http.Request) {
	validationGet(r)
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
	query := lib.QueryToMap(r.URL.Query())
	lib.SendJson(map[string]any{"msg": "Success", "page": query["page"], "limit": query["limit"]}, w)
}

func validationGet(r *http.Request) {
	query := r.URL.Query()
	for _, k := range []string{"page", "limit"} {
		if !query.Has(k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		panic(err)
	}
	if page < 0 {
		panic("bad: page must be positive or 0")
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		panic(err)
	}
	if limit < 1 {
		panic("bad: limit must be positive")
	}
	if query.Has("start") {
		layout := "2006-01-02"
		start, err := time.Parse(layout, query.Get("start"))
		if err != nil {
			panic(err)
		}
		if !start.Before(time.Now()) {
			panic("bad: start must be past or now")
		}
		if query.Has("end") {
			end, err := time.Parse(layout, query.Get("end"))
			if err != nil {
				panic(err)
			}
			if !end.Before(time.Now()) {
				panic("bad: end must be past or now")
			}
			if !end.After(start) {
				panic("bad: end must be after or same with start")
			}
		}
	}
}
