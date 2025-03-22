package repo

import (
	"database/sql"
	"log"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

type Wallet struct {
	Id      uint64  `json:"id"`
	Nm      string  `json:"nm"`
	Balance float64 `json:"balance"`
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
