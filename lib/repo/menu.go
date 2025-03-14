package repo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

type Menu struct {
	Id        uint64     `json:"id"`
	Label     string     `json:"label"`
	Link      *string    `json:"link"`
	Icon      *string    `json:"icon"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	ParentId  *uint64    `json:"parentId"`
	SubMenus  []Menu     `json:"subs"`
}

func AddMenu(tx *sql.Tx, menu Menu) error {
	mapArgs := map[string]any{"label": menu.Label}
	if menu.Link != nil {
		mapArgs["link"] = *menu.Link
	}
	if menu.Icon != nil {
		mapArgs["icon"] = *menu.Icon
	}
	keys := mapsutils.KeysOfMap(mapArgs)
	query := "INSERT INTO menus(" + strings.Join(keys, ",") + ") VALUES("
	query += strings.Join(arrayutils.Map(keys, func(_ string, k int) string { return fmt.Sprintf("$%d", k) }), ",") + ") RETURNING id"
	args := arrayutils.Map(keys, func(v string, _ int) any { return mapArgs[v] })
	log.Printf("Query \"%s\" with %v", query, args)
	rows, err := tx.Query(query, args...)
	defer func(rows *sql.Rows) {
		if rows != nil {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}
	}(rows)
	if err != nil {
		return err
	}
	var id uint64
	if rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("nothing added")
	}
	if menu.ParentId != nil {
		query = "MERGE INTO menu_has_menu m USING (SELECT $1 menu_id,$2 parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id "
		query += "WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id)"
		args = []any{id, *menu.ParentId}
		log.Printf("Query \"%s\" with %v", query, args)
		stmt, err := tx.Prepare(query)
		defer func(stmt *sql.Stmt) {
			if stmt != nil {
				if err := stmt.Close(); err != nil {
					panic(err)
				}
			}
		}(stmt)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(args...)
		if err != nil {
			return err
		}
	}
	return nil
}

func MenuById(tx *sql.Tx, id uint64) (Menu, error) {
	query := "SELECT m.id,m.label,m.link,m.icon,m.created_at,m.updated_at,mhm.parent_id FROM menus m "
	query += "LEFT JOIN menu_has_menu mhm ON mhm.menu_id=m.id WHERE m.id=$1 AND m.deleted_at IS NULL"
	return selectQueryAMenu(tx, query, id)
}

func MenuByRoleId(tx *sql.Tx, id uint64) ([]Menu, error) {
	query := "SELECT m.id,m.label,m.link,m.icon,m.created_at,m.updated_at,mhm.parent_id FROM role_has_menu rhm "
	query += "LEFT JOIN menus m ON m.id=rhm.menu_id AND m.deleted_at IS NULL LEFT JOIN menu_has_menu mhm ON mhm.menu_id=m.id "
	query += "WHERE rhm.role_id=$1 ORDER BY 1"
	return selectQueryMenus(tx, query, id)
}

func AllMenus(tx *sql.Tx) ([]Menu, error) {
	query := "SELECT m.id,m.label,m.link,m.icon,m.created_at,m.updated_at,mhm.parent_id FROM menus m "
	query += "LEFT JOIN menu_has_menu mhm ON mhm.menu_id=m.id WHERE m.deleted_at IS NULL"
	return selectQueryMenus2(tx, query)
}

func selectQueryMenus2(tx *sql.Tx, query string, args ...any) ([]Menu, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	rows, err := tx.Query(query, args...)
	defer func(rows *sql.Rows) {
		if rows != nil {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}
	}(rows)
	if err != nil {
		return nil, err
	}
	tmp, err := rowsToMenuTmp(rows)
	if err != nil {
		return nil, err
	}
	return arrayutils.Map(tmp, func(v map[string]any, _ int) Menu {
		mId := v["id"].(uint64)
		mLabel := v["label"].(string)
		mLink := v["link"].(*string)
		mIcon := v["icon"].(*string)
		mCreatedAt := v["createdAt"].(time.Time)
		mUpdatedAt := v["updatedAt"].(*time.Time)
		parentId := v["parentId"].(*uint64)
		return Menu{Id: mId, Label: mLabel, Link: mLink, Icon: mIcon, CreatedAt: mCreatedAt, UpdatedAt: mUpdatedAt, ParentId: parentId}
	}), nil
}

func selectQueryAMenu(tx *sql.Tx, query string, args ...any) (Menu, error) {
	results, err := selectQueryMenus(tx, query, args...)
	if err != nil {
		return Menu{}, err
	}
	if len(results) > 0 {
		return results[0], nil
	}
	return Menu{}, nil
}

func selectQueryMenus(tx *sql.Tx, query string, args ...any) ([]Menu, error) {
	log.Printf("Query \"%s\"", query)
	rows, err := tx.Query(query, args...)
	defer func(rows *sql.Rows) {
		if rows != nil {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}
	}(rows)
	if err != nil {
		return nil, err
	}
	tmp, err := rowsToMenuTmp(rows)
	if err != nil {
		return nil, err
	}
	menus := arrayutils.Grouping(tmp, func(v map[string]any, _ int) *uint64 { return v["parentId"].(*uint64) })
	parent := arrayutils.Map(menus[nil], func(v map[string]any, _ int) Menu { return tempRowToMap(v, menus) })
	arrayutils.Sort(parent, func(v1 Menu, v2 Menu) int {
		if v1.Id > v2.Id {
			return -1
		} else if v1.Id < v2.Id {
			return 1
		} else {
			return 0
		}
	})
	return parent, nil
}

func tempRowToMap(v map[string]any, menus map[*uint64][]map[string]any) Menu {
	id := v["id"].(uint64)
	keys := arrayutils.Filter(mapsutils.KeysOfMap(menus), func(v *uint64, _ int) bool { return v != nil && *v == id })
	return Menu{
		Id:        id,
		Label:     v["label"].(string),
		Link:      v["link"].(*string),
		Icon:      v["icon"].(*string),
		CreatedAt: v["createdAt"].(time.Time),
		UpdatedAt: v["updatedAt"].(*time.Time),
		ParentId:  nil,
		SubMenus: arrayutils.FlatMap(keys, func(v *uint64, _ int) []Menu {
			return arrayutils.Map(menus[v], func(v2 map[string]any, _ int) Menu { return tempRowToMap(v2, menus) })
		}),
	}
}

func rowsToMenuTmp(rows *sql.Rows) ([]map[string]any, error) {
	tmp := []map[string]any{}
	for rows.Next() {
		var mId uint64
		var mLabel string
		var mLink *string
		var mIcon *string
		var mCreatedAt time.Time
		var mUpdatedAt *time.Time
		var parentId *uint64
		err := rows.Scan(&mId, &mLabel, &mLink, &mIcon, &mCreatedAt, &mUpdatedAt, &parentId)
		if err != nil {
			return nil, err
		}
		line := map[string]any{
			"id":        mId,
			"label":     mLabel,
			"link":      mLink,
			"icon":      mIcon,
			"createdAt": mCreatedAt,
			"updatedAt": mUpdatedAt,
			"parentId":  parentId,
		}
		tmp = append(tmp, line)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tmp, nil
}
