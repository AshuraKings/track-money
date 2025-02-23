package repo

import (
	"net/http"
	"time"
	"track/lib"
	"track/lib/db"
)

type Role struct {
	Id        uint64     `json:"id"`
	Nm        string     `json:"nm"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func AllRoles(w http.ResponseWriter) ([]Role, bool) {
	var roles []Role
	clear := true
	db, err := db.DbConn()
	defer func() {
		clear = false
		lib.CloseDb(w, db)
	}()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id,nm,created_at,updated_at FROM roles WHERE deleted_at IS NULL")
	defer func() {
		clear = false
		lib.RowsClose(w, rows)
	}()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		role := Role{}
		if err = rows.Scan(&role.Id, &role.Nm, &role.CreatedAt, &role.UpdatedAt); err != nil {
			panic(err)
		}
		roles = append(roles, role)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return roles, clear
}
