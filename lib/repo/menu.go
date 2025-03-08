package repo

import (
	"database/sql"
	"log"
	"time"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

type Menu struct {
	Id        uint64
	Label     string
	Link      *string
	Icon      *string
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	ParentId  *uint64
	SubMenus  []Menu
}

func MenuById(tx *sql.Tx, id uint64) (Menu, error) {
	query := "SELECT m.id,m.label,m.link,m.icon,m.created_at,m.updated_at,mhm.parent_id FROM menus m "
	query += "LEFT JOIN menu_has_menu mhm ON mhm.menu_id=m.id "
	query += "WHERE m.id=$1 AND m.deleted_at IS NULL"
	return selectQueryAMenu(tx, query, id)
}

func MenuByRoleId(tx *sql.Tx, id uint64) ([]Menu, error) {
	query := "SELECT m.id,m.label,m.link,m.icon,m.created_at,m.updated_at,mhm.parent_id FROM role_has_menu rhm "
	query += "LEFT JOIN menus m ON m.id=rhm.menu_id AND m.deleted_at IS NULL LEFT JOIN menu_has_menu mhm ON mhm.menu_id=m.id "
	query += "WHERE rhm.role_id=$1"
	return selectQueryMenus(tx, query, id)
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
