package repo

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
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

func MapToMenu(body map[string]any) Menu {
	menu, keys := Menu{}, mapsutils.KeysOfMap(body)
	if arrayutils.Contains(keys, "id") {
		id := body["id"].(float64)
		if id < 1 {
			panic("bad: id not found")
		}
		menu.Id = uint64(id)
	}
	if !arrayutils.Contains(keys, "label") {
		panic("bad: label is required")
	}
	menu.Label = body["label"].(string)
	if arrayutils.Contains(keys, "link") {
		s := body["link"].(string)
		menu.Link = &s
	}
	if arrayutils.Contains(keys, "icon") {
		s := body["icon"].(string)
		menu.Icon = &s
	}
	if arrayutils.Contains(keys, "parentId") {
		id := body["parentId"].(float64)
		if id < 1 {
			panic("bad: parentId not found")
		}
		parent := uint64(id)
		menu.ParentId = &parent
	}
	return menu
}

func DelMenu(tx *sql.Tx, id uint64) error {
	menu, err := MenuById(tx, id)
	if err != nil {
		return err
	}
	if err = delASubMenu(tx, menu); err != nil {
		return err
	}
	query := "UPDATE menus SET deleted_at=now() WHERE id=$1"
	log.Printf("Query \"%s\" with %d", query, id)
	stmt, err := tx.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func EditMenu(tx *sql.Tx, menu Menu) error {
	if err := delASubMenu(tx, menu); err != nil {
		return err
	}
	if menu.ParentId != nil {
		if err := insASubMenu(tx, menu); err != nil {
			return err
		}
	}
	query, args, count := "UPDATE menus SET label=$1", []any{menu.Label}, 1
	if menu.Link != nil {
		count += 1
		query += fmt.Sprintf(",link=$%d", count)
		args = append(args, *menu.Link)
	}
	if menu.Icon != nil {
		count += 1
		query += fmt.Sprintf(",icon=$%d", count)
		args = append(args, *menu.Icon)
	}
	args = append(args, menu.Id)
	query += fmt.Sprintf(",updated_at=now() WHERE id=$%d", count+1)
	log.Printf("Query \"%s\" with %v", query, args)
	stmt, err := tx.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func insASubMenu(tx *sql.Tx, menu Menu) error {
	query := "MERGE INTO menu_has_menu m USING (SELECT $1::bigint menu_id,$2::bigint parent_id) AS n ON n.menu_id=m.menu_id AND n.parent_id=m.parent_id "
	query += "WHEN NOT MATCHED THEN INSERT(menu_id,parent_id) VALUES(n.menu_id,n.parent_id)"
	args := []any{menu.Id, *menu.ParentId}
	log.Printf("Query \"%s\" with %v", query, args)
	stmt, err := tx.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func delASubMenu(tx *sql.Tx, menu Menu) error {
	query := "DELETE FROM menu_has_menu WHERE menu_id=$1"
	log.Printf("Query \"%s\" with %d", query, menu.Id)
	stmt, err := tx.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(menu.Id)
	if err != nil {
		return err
	}
	return nil
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
	query += strings.Join(arrayutils.Map(keys, func(_ string, k int) string { return fmt.Sprintf("$%d", k+1) }), ",") + ") RETURNING id"
	args := arrayutils.Map(keys, func(v string, _ int) any { return mapArgs[v] })
	log.Printf("Query \"%s\" with %v", query, args)
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
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
		menu.Id = id
		if err = insASubMenu(tx, menu); err != nil {
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
	defer closeRows(rows)
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
	defer closeRows(rows)
	if err != nil {
		return nil, err
	}
	tmp, err := rowsToMenuTmp(rows)
	if err != nil {
		return nil, err
	}
	menus := arrayutils.Grouping(tmp, func(v map[string]any, _ int) *uint64 { return v["parentId"].(*uint64) })
	parent := arrayutils.Map(menus[nil], func(v map[string]any, _ int) Menu { return tempRowToMap(v, menus) })
	parentIds := arrayutils.Map(parent, func(v Menu, _ int) int { return int(v.Id) })
	sort.Ints(parentIds)
	return arrayutils.Map(arrayutils.Map(parentIds, func(v int, _ int) uint64 { return uint64(v) }), func(v uint64, _ int) Menu {
		summaries := arrayutils.Filter(parent, func(v2 Menu, _ int) bool { return v2.Id == v })
		if len(summaries) > 0 {
			return summaries[0]
		}
		return Menu{}
	}), nil
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

func closeRows(rows *sql.Rows) {
	if rows != nil {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}
}
