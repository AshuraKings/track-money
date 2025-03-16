package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

type UserWithRole struct {
	User User
	Role Role
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

func UserWithRoleByUserId(tx *sql.Tx, id uint64) (UserWithRole, error) {
	query := "SELECT u.id,u.nm,u.username,u.sandi,u.role_id,u.created_at,u.updated_at,r.id r_id,r.nm r_nm,r.created_at r_created_at,"
	query += "r.updated_at r_updated_at FROM users u LEFT JOIN roles r ON r.id=u.role_id AND r.deleted_at IS NULL WHERE u.deleted_at IS NULL AND u.id=$1 LIMIT 1"
	log.Printf("Query UserWithRoleByUserId : \"%s\"", query)
	return selectQueryAUserWithRole(tx, query, id)
}

func selectQueryAUserWithRole(tx *sql.Tx, query string, args ...any) (UserWithRole, error) {
	userWithRole := UserWithRole{}
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return UserWithRole{}, err
	}
	if rows.Next() {
		user, role := User{}, Role{}
		err = rows.Scan(&user.Id, &user.Nm, &user.Username, &user.Password, &user.RoleId, &user.CreatedAt, &user.UpdatedAt, &role.Id, &role.Nm, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return UserWithRole{}, err
		}
		userWithRole.Role = role
		userWithRole.User = user
	}
	if err = rows.Err(); err != nil {
		return UserWithRole{}, err
	}
	return userWithRole, nil
}

func selectQueryAUser(tx *sql.Tx, query string, args ...any) (User, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	user := User{Id: 0}
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return User{}, err
	}
	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.Nm, &user.Username, &user.Password, &user.RoleId, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return User{}, err
		}
	} else {
		return User{}, fmt.Errorf("nothing found")
	}
	if err = rows.Err(); err != nil {
		return User{}, err
	}
	return user, err
}

func selectQueryUsers(tx *sql.Tx, query string, args ...any) ([]User, error) {
	log.Printf("Query \"%s\"", query)
	var users []User
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
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
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Id)
	return err
}

func EditUser(tx *sql.Tx, user User) error {
	query := "UPDATE users SET nm=$1,username=$2,role_id=$3,updated_at=now() WHERE id=$4"
	stmt, err := tx.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Nm, user.Username, user.RoleId, user.Id)
	return err
}

func AddUser(tx *sql.Tx, user User) (uint64, error) {
	log.Println("Add User", user)
	err := addUser(tx, user)
	if err != nil {
		return 0, err
	}
	var id uint64
	rows, err := tx.Query("SELECT id FROM users WHERE username=$1 AND deleted_at IS NULL LIMIT 1", user.Username)
	defer closeRows(rows)
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
	log.Printf("Query \"%s\"", query)
	stmt, err := tx.Prepare(query)
	defer closeStmt(stmt)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.Nm, user.Password, user.Username, user.RoleId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Rows affected %d", count)
	if count != 0 {
		return errors.New("duplicate username")
	}
	return nil
}

func NewUser(nm string, username string, password string, role uint64) User {
	return User{Nm: nm, Username: username, Password: password, RoleId: role}
}
