package repo

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"track/lib/fieldbinding"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

type Transaksi struct {
	Kode         string   `json:"kode"`
	Ket          string   `json:"ket"`
	Amount       float64  `json:"amount"`
	Date         string   `json:"date"`
	AdminFee     float64  `json:"adminFee"`
	FromWallet   *Wallet  `json:"fromWallet"`
	ToWallet     *Wallet  `json:"toWallet"`
	ThisIncome   *Income  `json:"income"`
	ThisExpenses *Expense `json:"expense"`
	Doer         User     `json:"doer"`
}

func CountTransaksies(tx *sql.Tx, body map[string]any) (uint64, error) {
	keys := mapsutils.KeysOfMap(body)
	query, args := "SELECT count(1) FROM transaksi t WHERE t.deleted_at IS NULL ", []any{}
	if arrayutils.Contains(keys, "ket") && body["ket"] != nil {
		ket := body["ket"].(string)
		query += fmt.Sprintf("AND t.ket ILIKE $%d ", len(args)+1)
		args = append(args, "%"+ket+"%")
	}
	if arrayutils.Contains(keys, "start") && body["start"] != nil {
		start := body["start"].(time.Time)
		if arrayutils.Contains(keys, "end") && body["end"] != nil {
			end := body["end"].(time.Time)
			query += fmt.Sprintf("AND t.trx_date BETWEEN $%d ", len(args)+1)
			args = append(args, start)
			query += fmt.Sprintf("AND $%d ", len(args)+1)
			args = append(args, end)
		} else {
			query += fmt.Sprintf("AND t.trx_date>=$%d ", len(args)+1)
			args = append(args, start)
		}
	}
	return selectCount(tx, query, args...)
}

func AllTransaksies(tx *sql.Tx, body map[string]any) ([]Transaksi, error) {
	keys := mapsutils.KeysOfMap(body)
	query, args := "SELECT t.kode,t.ket,t.amount,t.trx_date,t.admin_fee,t.from_wallet_id,fw.nm fw_nm,fw.balance fw_balance,", []any{}
	query += "t.to_wallet_id,tw.nm tw_nm,tw.balance tw_balance,t.income_id,i.nm i_nm,t.expenses_id,e.nm e_nm,t.doer_id,u.nm u_nm,u.username u_username,"
	query += "u.role_id u_role_id,u.created_at u_created_at,u.updated_at u_updated_at "
	query += "FROM transaksi t LEFT JOIN wallets fw ON t.from_wallet_id=fw.id LEFT JOIN wallets tw ON t.to_wallet_id=tw.id "
	query += "LEFT JOIN incomes i ON i.id=t.income_id LEFT JOIN expenses e ON e.id=t.expenses_id LEFT JOIN users u ON u.id=t.doer_id "
	query += "WHERE t.deleted_at IS NULL "
	if arrayutils.Contains(keys, "ket") && body["ket"] != nil {
		ket := body["ket"].(string)
		query += fmt.Sprintf("AND t.ket ILIKE $%d ", len(args)+1)
		args = append(args, "%"+ket+"%")
	}
	if arrayutils.Contains(keys, "start") && body["start"] != nil {
		start := body["start"].(time.Time)
		if arrayutils.Contains(keys, "end") && body["end"] != nil {
			end := body["end"].(time.Time)
			query += fmt.Sprintf("AND t.trx_date BETWEEN $%d ", len(args)+1)
			args = append(args, start)
			query += fmt.Sprintf("AND $%d ", len(args)+1)
			args = append(args, end)
		} else {
			query += fmt.Sprintf("AND t.trx_date>=$%d ", len(args)+1)
			args = append(args, start)
		}
	}
	query += fmt.Sprintf("ORDER BY t.trx_date DESC LIMIT $%d ", len(args)+1)
	args = append(args, body["limit"])
	query += fmt.Sprintf("OFFSET $%d ", len(args)+1)
	args = append(args, body["page"])
	return selectTransaksies(tx, query, args...)
}

func selectCount(tx *sql.Tx, query string, args ...any) (uint64, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return 0, err
	}
	rowsMap, err := rowsToMap(rows)
	if err != nil {
		return 0, err
	}
	if len(rowsMap) != 0 {
		row := rowsMap[0]
		count1 := row["count"].(int64)
		return uint64(count1), nil
	}
	return 0, fmt.Errorf("count query error")
}

func selectTransaksies(tx *sql.Tx, query string, args ...any) ([]Transaksi, error) {
	log.Printf("Query \"%s\" with %v", query, args)
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return nil, err
	}
	rowsMap, err := rowsToMap(rows)
	if err != nil {
		return nil, err
	}
	return arrayutils.Map(rowsMap, mapToTransaksi), nil
}

func mapToTransaksi(v map[string]any, _ int) Transaksi {
	t := Transaksi{}
	t.Kode = v["kode"].(string)
	t.Ket = v["ket"].(string)
	t.AdminFee = v["admin_fee"].(float64)
	t.Amount = v["amount"].(float64)
	date := v["trx_date"].(time.Time)
	t.Date = date.Format("2006-01-02")
	id, roleId := v["doer_id"].(int64), v["u_role_id"].(int64)
	t.Doer.Id = uint64(id)
	t.Doer.Nm = v["u_nm"].(string)
	t.Doer.Username = v["u_username"].(string)
	t.Doer.RoleId = uint64(roleId)
	t.Doer.CreatedAt = v["u_created_at"].(time.Time)
	if v["u_updated_at"] != nil {
		date1 := v["u_updated_at"].(time.Time)
		t.Doer.UpdatedAt = &date1
	}
	if v["expenses_id"] != nil {
		e := Expense{}
		id := v["expenses_id"].(int64)
		e.Id = uint64(id)
		e.Nm = v["e_nm"].(string)
		t.ThisExpenses = &e
	}
	if v["income_id"] != nil {
		i := Income{}
		id := v["income_id"].(int64)
		i.Id = uint64(id)
		i.Nm = v["i_nm"].(string)
		t.ThisIncome = &i
	}
	if v["from_wallet_id"] != nil {
		fw := Wallet{}
		id := v["from_wallet_id"].(int64)
		fw.Id = uint64(id)
		fw.Nm = v["fw_nm"].(string)
		fw.Balance = v["fw_balance"].(float64)
		t.FromWallet = &fw
	}
	if v["to_wallet_id"] != nil {
		fw := Wallet{}
		id := v["to_wallet_id"].(int64)
		fw.Id = uint64(id)
		fw.Nm = v["tw_nm"].(string)
		fw.Balance = v["tw_balance"].(float64)
		t.ToWallet = &fw
	}
	return t
}

func rowsToMap(rows *sql.Rows) ([]map[string]any, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	log.Printf("Cols %v", cols)
	fb := fieldbinding.NewFieldBinding()
	fb.PutFields(cols)
	res := []map[string]any{}
	for rows.Next() {
		row := map[string]any{}
		if err = rows.Scan(fb.GetFieldPtrArr()...); err != nil {
			return nil, err
		}
		for _, col := range cols {
			row[col] = fb.Get(col)
		}
		res = append(res, row)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
