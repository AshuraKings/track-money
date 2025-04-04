package repo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

type Wallet struct {
	Id      uint64  `json:"id"`
	Nm      string  `json:"nm"`
	Balance float64 `json:"balance"`
}

func DelWallet(tx *sql.Tx, id uint64) error {
	query, args := "UPDATE wallets SET deleted_at=now() WHERE id=$1", []any{id}
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

func UpdateWalletBalance(tx *sql.Tx, id uint64, change float64) error {
	query, args := "MERGE INTO wallets w USING (SELECT $1::bigint id,$2::numeric change) AS c ON w.id=c.id AND w.deleted_at IS NULL ", []any{id, change}
	query += "WHEN MATCHED THEN UPDATE SET updated_at=now(),balance=balance+c.change"
	return stmtExec(tx, query, args...)
}

func EditWallet(tx *sql.Tx, id uint64, mapChange map[string]any) error {
	keys, constKeys := mapsutils.KeysOfMap(mapChange), []string{"nm", "balanceUp", "balanceDown"}
	if len(keys) < 1 || arrayutils.AllOf(keys, func(v string, _ int) bool { return !arrayutils.Contains(constKeys, v) }) {
		return fmt.Errorf("bad: changes not found")
	}
	query, args, queries := "UPDATE wallets SET ", []any{}, []string{}
	if arrayutils.Contains(keys, "nm") {
		nm := mapChange["nm"].(string)
		queries = append(queries, fmt.Sprintf("nm=$%d", (len(args)+1)))
		args = append(args, nm)
	}
	if arrayutils.Contains(keys, "balanceUp") {
		balanceUp := mapChange["balanceUp"].(float64)
		queries = append(queries, fmt.Sprintf("balance=balance+$%d", (len(args)+1)))
		args = append(args, balanceUp)
	}
	if arrayutils.Contains(keys, "balanceDown") {
		balanceDown := mapChange["balanceDown"].(float64)
		queries = append(queries, fmt.Sprintf("balance=balance-$%d", (len(args)+1)))
		args = append(args, balanceDown)
	}
	query += strings.Join(queries, ",")
	query += fmt.Sprintf(",updated_at=now() WHERE id=$%d", (len(args) + 1))
	args = append(args, id)
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

func AddWallet(tx *sql.Tx, wallet Wallet) error {
	query := "MERGE INTO wallets w USING(SELECT $1 nm,$2::numeric balance) AS n ON w.nm=n.nm WHEN NOT MATCHED THEN INSERT(nm,balance) VALUES(n.nm,n.balance)"
	args := []any{wallet.Nm, wallet.Balance}
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

func GetWallets(tx *sql.Tx, id uint64) ([]Wallet, error) {
	query, args := "SELECT id,nm,balance FROM wallets WHERE deleted_at IS NULL AND id=$1", []any{id}
	return selectQueryWallets(tx, query, args...)
}

func AllWallet(tx *sql.Tx) ([]Wallet, error) {
	query := "SELECT id,nm,balance FROM wallets WHERE deleted_at IS NULL"
	return selectQueryWallets(tx, query)
}

func selectQueryWallets(tx *sql.Tx, query string, args ...any) ([]Wallet, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	var results []Wallet
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var wallet Wallet
		if err = rows.Scan(&wallet.Id, &wallet.Nm, &wallet.Balance); err != nil {
			return nil, err
		}
		results = append(results, wallet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func FromMapToWallet(body map[string]any) Wallet {
	w := Wallet{Nm: body["nm"].(string), Balance: body["balance"].(float64)}
	keys := mapsutils.KeysOfMap(body)
	if arrayutils.Contains(keys, "id") {
		id := body["id"].(float64)
		w.Id = uint64(id)
	}
	return w
}
