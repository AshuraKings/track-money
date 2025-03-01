package repo

import (
	"database/sql"
	"time"
)

type User struct {
	Id        uint64
	Nm        string
	Username  string
	Password  string
	RoleId    uint64
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func UserByUsername(tx *sql.Tx, username string) (User, error) {
	query := "SELECT id,nm,username,sandi,role_id,created_at,updated_at FROM users WHERE deleted_at IS NULL AND username=$1 LIMIT 1"
	return selectQueryAUser(tx, query, username)
}

func UserById(tx *sql.Tx, id uint64) (User, error) {
	query := "SELECT id,nm,username,sandi,role_id,created_at,updated_at FROM users WHERE deleted_at IS NULL AND id=$1 LIMIT 1"
	return selectQueryAUser(tx, query, id)
}

func AllUsers(tx *sql.Tx) ([]User, error) {
	query := "SELECT id,nm,username,sandi,role_id,created_at,updated_at FROM users WHERE deleted_at IS NULL"
	return selectQueryUsers(tx, query)
}

func selectQueryAUser(tx *sql.Tx, query string, args ...any) (User, error) {
	user := User{Id: 0}
	rows, err := tx.Query(query, args...)
	defer func(rows *sql.Rows) {
		if rows != nil {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}
	}(rows)
	if err != nil {
		return User{}, err
	}
	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.Nm, &user.Username, &user.Password, &user.RoleId, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return User{}, err
		}
	}
	if err = rows.Err(); err != nil {
		return User{}, err
	}
	return user, err
}

func selectQueryUsers(tx *sql.Tx, query string, args ...any) ([]User, error) {
	var users []User
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
		user := User{}
		if err = rows.Scan(&user.Id, &user.Nm, &user.Username, &user.Password, &user.RoleId, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}

func DeleteUser(tx *sql.Tx, user User) error {
	query := "UPDATE users SET deleted_at=now() WHERE id=$1"
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
	_, err = stmt.Exec(user.Id)
	return err
}

func EditUser(tx *sql.Tx, user User) error {
	query := "UPDATE users SET nm=$1,username=$2,updated_at=now() WHERE id=$3"
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
	_, err = stmt.Exec(user.Nm, user.Username, user.Id)
	return err
}

func AddUser(tx *sql.Tx, user User) (uint64, error) {
	err := addUser(tx, user)
	if err != nil {
		return 0, err
	}
	var id uint64
	rows, err := tx.Query("SELECT id FROM users WHERE username=$1 AND deleted_at IS NULL LIMIT 1", user.Username)
	defer func(rows *sql.Rows) {
		if rows != nil {
			if err := rows.Close(); err != nil {
				panic(err)
			}
		}
	}(rows)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, nil
		}
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}
	return id, err
}

func addUser(tx *sql.Tx, user User) error {
	query := "MERGE INTO users u USING(SELECT $1 nm,$2 password,$3 username,$4::bigint role_id) AS n ON u.username=n.username "
	query = query + "WHEN NOT MATCHED THEN INSERT(nm,sandi,username,role_id) VALUES(n.nm,n.password,n.username,n.role_id)"
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
	_, err = stmt.Exec(user.Nm, user.Password, user.Username, user.RoleId)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(nm string, username string, password string, role uint64) User {
	return User{Nm: nm, Username: username, Password: password, RoleId: role}
}
