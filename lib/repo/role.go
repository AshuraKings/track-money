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

func DelRole(tx *sql.Tx, role Role) error {
	query := "UPDATE roles SET deleted_at=now() WHERE id=$1"
	log.Printf("Query \"%s\" with \"%v\"", query, role)
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
	_, err = stmt.Exec(role.Id)
	if err != nil {
		return err
	}
	return nil
}

func EditRole(tx *sql.Tx, role Role) error {
	query := "UPDATE roles SET nm=$1,updated_at=now() WHERE id=$2"
	log.Printf("Query \"%s\" with \"%v\"", query, role)
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
	_, err = stmt.Exec(role.Nm, role.Id)
	if err != nil {
		return err
	}
	return nil
}

func AddRole(tx *sql.Tx, name string) error {
	query := "MERGE INTO roles r USING(SELECT $1 nm) AS n ON r.nm=n.nm WHEN NOT MATCHED THEN INSERT(nm) VALUES(n.nm)"
	log.Printf("Query \"%s\" with \"%s\"", query, name)
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
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
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
