package repo

import (
	"database/sql"
	"log"
)

type Wallet struct {
	Id      uint64  `json:"id"`
	Nm      string  `json:"nm"`
	Balance float64 `json:"balance"`
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
