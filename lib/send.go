package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
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

func TxClose(tx *sql.Tx, w http.ResponseWriter) {
	if tx != nil {
		if r := recover(); r != nil {
			msg := fmt.Sprint("", r)
			log.Println(msg)
			log.Println(string(debug.Stack()))
			if strings.HasPrefix(msg, "method") {
				w.WriteHeader(405)
			} else if strings.HasPrefix(msg, "Token") {
				w.WriteHeader(403)
			} else if strings.HasPrefix(msg, "bad:") {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			if err := tx.Rollback(); err != nil {
				panic(err)
			}
			SendJson(map[string]string{"msg": msg}, w)
		} else {
			if err := tx.Commit(); err != nil {
				panic(err)
			}
		}
	} else {
		DefaultError(w)
	}
}

func DefaultError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		msg := fmt.Sprint("", r)
		log.Println(msg)
		log.Println(string(debug.Stack()))
		if strings.HasPrefix(msg, "method") {
			w.WriteHeader(405)
		} else if strings.HasPrefix(msg, "Token") {
			w.WriteHeader(403)
		} else if strings.HasPrefix(msg, "bad:") {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		SendJson(map[string]string{"msg": msg}, w)
	}
}

func SendJson(resp any, w http.ResponseWriter) {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
