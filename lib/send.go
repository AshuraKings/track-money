package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func StmtClose(w http.ResponseWriter, stmt *sql.Stmt) {
	DefaultError(w)
	if stmt != nil {
		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}
}

func RowsClose(w http.ResponseWriter, rows *sql.Rows) {
	DefaultError(w)
	if rows != nil {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}
}

func CloseDb(w http.ResponseWriter, db *sql.DB) {
	DefaultError(w)
	if db != nil {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}
}

func DefaultError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		w.WriteHeader(500)
		msg := fmt.Sprint("", r)
		log.Println(msg)
		log.Println(string(debug.Stack()))
		SendJson(map[string]string{"msg": msg}, w)
	}
}

func SendJson(resp any, w http.ResponseWriter) {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
}
