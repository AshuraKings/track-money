package repo

import (
	"database/sql"
	"log"
)

type Income struct {
	Id uint64 `json:"id"`
	Nm string `json:"nm"`
}

func DelIncome(tx *sql.Tx, id uint64) error {
	query, args := "UPDATE incomes SET deleted_at=now() WHERE id=$1", []any{id}
	return stmtExec(tx, query, args...)
}

func AddIncome(tx *sql.Tx, nm string) error {
	query, args := "MERGE INTO incomes i USING (SELECT $1 nm) AS n ON i.nm=n.nm WHEN NOT MATCHED THEN INSERT (nm) VALUES(n.nm)", []any{nm}
	return stmtExec(tx, query, args...)
}

func stmtExec(tx *sql.Tx, query string, args ...any) error {
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
