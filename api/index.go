package handler

import (
	"errors"
	"log"
	"net/http"
	"os"
	"track/lib"
	"track/lib/authed"
	"track/lib/authed/expenses"
	"track/lib/authed/incomes"
	menus1 "track/lib/authed/menus"
	"track/lib/authed/refresh"
	"track/lib/authed/roles"
	"track/lib/authed/roles/menus"
	"track/lib/authed/transactions"
	"track/lib/authed/users"
	"track/lib/authed/wallets"
	"track/lib/db"
	"track/lib/login"
	"track/lib/logout"
	"track/lib/register"
	"track/lib/repo"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	os.Setenv("TZ", "Asia/Jakarta")
	header := r.Header
	if header.Get("ai-path") != "" {
		myPath := header.Get("ai-path")
		log.Printf("[%s] /api%s", r.Method, myPath)
		if myPath == "/authed" {
			authed.Handler(w, r)
		} else if myPath == "/authed/wallets" {
			wallets.Handler(w, r)
		} else if myPath == "/authed/users" {
			users.Handler(w, r)
		} else if myPath == "/authed/transactions" {
			transactions.Handler(w, r)
		} else if myPath == "/authed/roles/menus" {
			menus.Handler(w, r)
		} else if myPath == "/authed/roles" {
			roles.Handler(w, r)
		} else if myPath == "/authed/menus" {
			menus1.Handler(w, r)
		} else if myPath == "/authed/incomes" {
			incomes.Handler(w, r)
		} else if myPath == "/authed/refresh" {
			refresh.Handler(w, r)
		} else if myPath == "/authed/expenses" {
			expenses.Handler(w, r)
		} else if myPath == "/register" {
			register.Handler(w, r)
		} else if myPath == "/logout" {
			logout.Handler(w, r)
		} else if myPath == "/login" {
			login.Handler(w, r)
		} else {
			panic(errors.New("method not allowed"))
		}
	} else {
		if r.Method == "GET" {
			resp := map[string]any{}
			resp["msg"] = "Hello World"
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
			resp["roles"] = roles
			lib.SendJson(resp, w)
		} else {
			panic(errors.New("method not allowed"))
		}
	}
}
