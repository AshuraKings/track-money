package repo

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
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

func AddTransksi(tx *sql.Tx, body map[string]any) error {
	log.Printf("Value %v", body)
	keys := mapsutils.KeysOfMap(body)
	query, args, cols := "MERGE INTO transaksi t USING (SELECT ", []any{}, []string{}
	kode := body["kode"].(string)
	args = append(args, kode)
	query += fmt.Sprintf("$%d kode,", len(args))
	cols = append(cols, "kode")
	query += fmt.Sprintf("$%d ket,", len(args)+1)
	args = append(args, body["ket"].(string))
	cols = append(cols, "ket")
	query += fmt.Sprintf("$%d::numeric amount,", len(args)+1)
	args = append(args, body["amount"].(float64))
	cols = append(cols, "amount")
	args = append(args, body["date"].(time.Time))
	query += fmt.Sprintf("$%d::date trx_date,", len(args))
	cols = append(cols, "trx_date")
	args = append(args, body["admin"].(float64))
	query += fmt.Sprintf("$%d::numeric admin_fee,", len(args))
	cols = append(cols, "admin_fee")
	if arrayutils.Contains(keys, "fw") {
		args = append(args, body["fw"].(float64))
		query += fmt.Sprintf("$%d::bigint from_wallet_id,", len(args))
		cols = append(cols, "from_wallet_id")
	}
	if arrayutils.Contains(keys, "income") {
		args = append(args, body["income"].(float64))
		query += fmt.Sprintf("$%d::bigint income_id,", len(args))
		cols = append(cols, "income_id")
	}
	if arrayutils.Contains(keys, "tw") {
		args = append(args, body["tw"].(float64))
		query += fmt.Sprintf("$%d::bigint to_wallet_id,", len(args))
		cols = append(cols, "to_wallet_id")
	}
	if arrayutils.Contains(keys, "expense") {
		args = append(args, body["expense"].(float64))
		query += fmt.Sprintf("$%d::bigint expenses_id,", len(args))
		cols = append(cols, "expenses_id")
	}
	args = append(args, body["doer"].(int))
	query += fmt.Sprintf("$%d::bigint doer_id)", len(args))
	cols = append(cols, "doer_id")
	query += " AS n ON t.kode=n.kode WHEN NOT MATCHED THEN INSERT"
	query += fmt.Sprintf("(%s) ", strings.Join(cols, ","))
	query += fmt.Sprintf("VALUES(%s)", strings.Join(arrayutils.Map(cols, func(v string, _ int) string { return "n." + v }), ","))
	return stmtExec(tx, query, args...)
}

func GenKodeTransaksi() string {
	prefix, additional := time.Now().Format("0601"), randomString(16)
	return prefix + additional
}

func randomString(length int) string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
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

func FirstDateTransaksi(tx *sql.Tx) (string, error) {
	query, args := "SELECT MIN(t.trx_date) pertama FROM transaksi t", []any{}
	log.Printf("Query \"%s\" with %v", query, args)
	rows, err := tx.Query(query, args...)
	defer closeRows(rows)
	if err != nil {
		return "", err
	}
	rowsMap, err := rowsToMap(rows)
	if err != nil {
		return "", err
	}
	if len(rowsMap) == 0 || rowsMap[0]["pertama"] == nil {
		return time.Now().Format("2006-01-02"), nil
	}
	return rowsMap[0]["pertama"].(time.Time).Format("2006-01-02"), nil
}

func AllTransaksies(tx *sql.Tx, body map[string]any) ([]Transaksi, error) {
	keys := mapsutils.KeysOfMap(body)
	query, args := "SELECT t.kode,t.ket,t.amount::text amount,t.trx_date,t.admin_fee::text admin_fee,t.from_wallet_id,fw.nm fw_nm,fw.balance::text fw_balance,", []any{}
	query += "t.to_wallet_id,tw.nm tw_nm,tw.balance::text tw_balance,t.income_id,i.nm i_nm,t.expenses_id,e.nm e_nm,t.doer_id,u.nm u_nm,u.username u_username,"
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
	adminFee1 := v["admin_fee"].(string)
	adminFee, err := strconv.ParseFloat(adminFee1, 64)
	t.AdminFee = adminFee
	if err != nil {
		panic(err)
	}
	amount1 := v["amount"].(string)
	t.Amount, err = strconv.ParseFloat(amount1, 64)
	if err != nil {
		panic(err)
	}
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
		balance := v["fw_balance"].(string)
		fw.Balance, err = strconv.ParseFloat(balance, 64)
		if err != nil {
			panic(err)
		}
		t.FromWallet = &fw
	}
	if v["to_wallet_id"] != nil {
		fw := Wallet{}
		id := v["to_wallet_id"].(int64)
		fw.Id = uint64(id)
		fw.Nm = v["tw_nm"].(string)
		balance := v["tw_balance"].(string)
		fw.Balance, err = strconv.ParseFloat(balance, 64)
		if err != nil {
			panic(err)
		}
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
