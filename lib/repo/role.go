package repo

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Role struct {
	Id        uint64     `json:"id"`
	Nm        string     `json:"nm"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func AllRoles(tx *sql.Tx) ([]Role, error) {
	query := "SELECT id,nm,created_at,updated_at FROM roles WHERE deleted_at IS NULL"
	return selectQueryRoles(tx, query)
}

func RoleById(tx *sql.Tx, id uint64) (Role, error) {
	query := "SELECT id,nm,created_at,updated_at FROM roles WHERE deleted_at IS NULL AND id=$1 LIMIT 1"
	return selectQueryARole(tx, query, id)
}

func selectQueryARole(tx *sql.Tx, query string, args ...any) (Role, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	role := Role{}
	rows, err := tx.Query(query, args...)
	defer func(rows *sql.Rows) {
		if rows != nil {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}
	}(rows)
	if err != nil {
		return Role{}, err
	}
	if rows.Next() {
		if err = rows.Scan(&role.Id, &role.Nm, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return Role{}, err
		}
	} else {
		return Role{}, errors.New("role not found")
	}
	if err = rows.Err(); err != nil {
		return Role{}, err
	}
	return role, nil
}

func selectQueryRoles(tx *sql.Tx, query string, args ...any) ([]Role, error) {
	log.Printf("Query \"%s\"", query)
	var roles []Role
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
	for rows.Next() {
		role := Role{}
		if err = rows.Scan(&role.Id, &role.Nm, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}
