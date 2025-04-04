package transactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultError(w)
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin", "fin"}) {
			return
		}
		getting(w, r)
	} else if r.Method == "POST" {
		if !session.ValidationRole(w, r, []string{"admin", "fin"}) {
			return
		}
		posting(w, r)
	} else if r.Method == "DELETE" {
		if !session.ValidationRole(w, r, []string{"admin"}) {
			return
		}
		lib.SendJson(map[string]any{"msg": "Success"}, w)
	} else {
		panic("method not allowed")
	}
}

func posting(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		panic(err)
	}
	validationPost(body)
	db, err := db.DbConn()
	defer lib.CloseDb(w, db)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	defer lib.TxClose(tx, w)
	if err != nil {
		panic(err)
	}
	body["kode"] = repo.GenKodeTransaksi()
	date, err := time.Parse("2006-01-02", body["date"].(string))
	if err != nil {
		panic(err)
	}
	body["date"] = date
	claims := session.ParseToken(r)
	sub := claims["sub"].(int64)
	body["doer"] = sub
	if err = repo.AddTransksi(tx, body); err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationPost(body map[string]any) {
	keys := mapsutils.KeysOfMap(body)
	for _, k := range []string{"admin", "amount", "date", "ket"} {
		if !arrayutils.Contains(keys, k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	hasFw, hasIncome := arrayutils.Contains(keys, "fw"), arrayutils.Contains(keys, "income")
	if hasFw && hasIncome {
		panic("bad: wallet and income are required 1 only")
	}
	if !hasFw && !hasIncome {
		panic("bad: wallet and income are conflicted")
	}
	if hasFw {
		fw := body["fw"].(float64)
		if fw <= 0 {
			panic("bad: from wallet is not found")
		}
	}
	if hasIncome {
		income := body["income"].(float64)
		if income <= 0 {
			panic("bad: income is not found")
		}
	}
	hasTw, hasExpense := arrayutils.Contains(keys, "tw"), arrayutils.Contains(keys, "expense")
	if !hasTw && !hasExpense {
		panic("bad: wallet and expense are required 1 only")
	}
	if hasTw && hasExpense {
		panic("bad: wallet and expense are conflicted")
	}
	if hasTw {
		tw := body["tw"].(float64)
		if tw <= 0 {
			panic("bad: to wallet is not found")
		}
	}
	if hasExpense {
		expense := body["expense"].(float64)
		if expense <= 0 {
			panic("bad: expense is not found")
		}
	}
	ket := body["ket"].(string)
	if ket == "" {
		panic("bad: ket must be not empty")
	}
	date1 := body["date"].(string)
	if date1 == "" {
		panic("bad: invalid date")
	}
	layout := "2006-01-02"
	date, err := time.Parse(layout, date1)
	if err != nil {
		panic(err)
	}
	if date.After(time.Now()) {
		panic("bad: date must be before or now")
	}
	amount := body["amount"].(float64)
	if amount <= 0 {
		panic("bad: amount must be positive")
	}
	admin := body["admin"].(float64)
	if admin < 0 {
		panic("bad: adminFee must be positive or zero")
	}
}

func getting(w http.ResponseWriter, r *http.Request) {
	validationGet(r)
	query := lib.QueryToMap(r.URL.Query())
	db, err := db.DbConn()
	defer lib.CloseDb(w, db)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	defer lib.TxClose(tx, w)
	if err != nil {
		panic(err)
	}
	where := mapsutils.Map(query, mapConvert1)
	transaksies, err := repo.AllTransaksies(tx, where)
	if err != nil {
		panic(err)
	}
	count, err := repo.CountTransaksies(tx, where)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{
		"msg":         "Success",
		"page":        query["page"],
		"limit":       query["limit"],
		"transaksies": transaksies,
		"count":       count,
	}, w)
}

func mapConvert1(v, k string) any {
	if arrayutils.Contains([]string{"limit", "page"}, k) {
		r, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return r
	}
	if arrayutils.Contains([]string{"start", "end"}, k) {
		r, err := time.Parse("2006-01-02", v)
		if err != nil {
			panic(err)
		}
		return r
	}
	return v
}

func validationGet(r *http.Request) {
	query := r.URL.Query()
	for _, k := range []string{"page", "limit"} {
		if !query.Has(k) {
			panic(fmt.Sprintf("bad: %s is required", k))
		}
	}
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		panic(err)
	}
	if page < 0 {
		panic("bad: page must be positive or 0")
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		panic(err)
	}
	if limit < 1 {
		panic("bad: limit must be positive")
	}
	if query.Has("start") {
		layout := "2006-01-02"
		start, err := time.Parse(layout, query.Get("start"))
		if err != nil {
			panic(err)
		}
		if !start.Before(time.Now()) {
			panic("bad: start must be past or now")
		}
		if query.Has("end") {
			end, err := time.Parse(layout, query.Get("end"))
			if err != nil {
				panic(err)
			}
			if !end.Before(time.Now()) {
				panic("bad: end must be past or now")
			}
			if !end.After(start) {
				panic("bad: end must be after or same with start")
			}
		}
	}
}
