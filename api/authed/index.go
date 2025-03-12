package handler

import (
	"net/http"
	"sort"
	"strconv"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"

	arrayutils "github.com/AchmadRifai/array-utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{}) {
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
		menus, err := repo.MenuByRoleId(tx, user.Role.Id)
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
			"menus": arrayutils.Map(menus, menuToMap),
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

func menuToMap(v repo.Menu, _ int) map[string]any {
	subs := v.SubMenus
	indexed := arrayutils.Map(subs, func(v repo.Menu, _ int) int { return int(v.Id) })
	sort.Ints(indexed)
	return map[string]any{
		"createdAt": v.CreatedAt,
		"icon":      v.Icon,
		"id":        v.Id,
		"label":     v.Label,
		"link":      v.Link,
		"updatedAt": v.UpdatedAt,
		"parent":    v.ParentId,
		"subs": arrayutils.Map(arrayutils.Map(indexed, func(v int, _ int) repo.Menu {
			for _, m := range subs {
				if m.Id == uint64(v) {
					return m
				}
			}
			return repo.Menu{}
		}), menuToMap),
	}
}
