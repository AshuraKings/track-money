package handler

import (
	"net/http"
	"strconv"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin", "viewer-out", "viewer", "fin"}) {
			return
		}
		claim := session.ParseToken(r)
		id, err := strconv.ParseUint(claim["sub"].(string), 10, 64)
		if err != nil {
			panic(err)
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
		user, err := repo.UserWithRoleByUserId(tx, id)
		if err != nil {
			panic(err)
		}
		lib.SendJson(map[string]any{
			"role": map[string]any{
				"id":        user.Role.Id,
				"nm":        user.Role.Nm,
				"createdAt": user.Role.CreatedAt,
				"updatedAt": user.Role.UpdatedAt,
			},
			"user": map[string]any{
				"id":        user.User.Id,
				"nm":        user.User.Nm,
				"username":  user.User.Username,
				"roleId":    user.User.RoleId,
				"createdAt": user.User.CreatedAt,
				"updatedAt": user.User.UpdatedAt,
			},
			"id": id,
		}, w)
	} else {
		panic("method not allowed")
	}
}
