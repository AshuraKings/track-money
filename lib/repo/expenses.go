package repo

import (
	"database/sql"
	"log"
)

type Expense struct {
	Id uint64 `json:"id"`
	Nm string `json:"nm"`
}

func DelExpense(tx *sql.Tx, id uint64) error {
	query, args := "UPDATE expenses SET deleted_at=now() WHERE id=$1", []any{id}
	return stmtExec(tx, query, args...)
}

func AddExpense(tx *sql.Tx, nm string) error {
	query, args := "MERGE INTO expenses e USING (SELECT $1 nm) AS n ON e.nm=n.nm WHEN NOT MATCHED THEN INSERT (nm) VALUES(n.nm)", []any{nm}
	return stmtExec(tx, query, args...)
}

func AllExpenses(tx *sql.Tx) ([]Expense, error) {
	query := "SELECT id,nm FROM expenses WHERE deleted_at IS NULL"
	return selectQueryExpenses(tx, query)
}

func selectQueryExpenses(tx *sql.Tx, query string, args ...any) ([]Expense, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	var expenses []Expense
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var expense Expense
		err = rows.Scan(&expense.Id, &expense.Nm)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return expenses, nil
}
