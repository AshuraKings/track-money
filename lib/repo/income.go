package repo

import (
	"database/sql"
	"log"
)

type Income struct {
	Id uint64 `json:"id"`
	Nm string `json:"nm"`
}

func AllIncome(tx *sql.Tx) ([]Income, error) {
	query := "SELECT id,nm FROM incomes WHERE deleted_at IS NULL"
	return selectQueryIncomes(tx, query)
}

func selectQueryIncomes(tx *sql.Tx, query string, args ...any) ([]Income, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	var incomes []Income
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var income Income
		err = rows.Scan(&income.Id, &income.Nm)
		if err != nil {
			return nil, err
		}
		incomes = append(incomes, income)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return incomes, nil
}
