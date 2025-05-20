package monthly

import (
	"net/http"
	"time"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"
	"track/lib/session"

	arrayutils "github.com/AchmadRifai/array-utils"
	mapsutils "github.com/AchmadRifai/maps-utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if !session.ValidationRole(w, r, []string{"admin", "fin", "viewer"}) {
			return
		}
		query := lib.QueryToMap(r.URL.Query())
		if len(query) > 0 {
			getting(w, query)
		} else {
			gettingFirstDate(w)
		}
	} else {
		panic("method not allowed")
	}
}

func getting(w http.ResponseWriter, query map[string]string) {
	validationGet(query)
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
	lib.SendJson(map[string]any{"msg": "Success"}, w)
}

func validationGet(query map[string]string) {
	if !arrayutils.Contains(mapsutils.KeysOfMap(query), "month") {
		panic("month not found")
	}
	_, err := time.Parse("2006-02", query["month"])
	if err != nil {
		panic(err)
	}
}

func gettingFirstDate(w http.ResponseWriter) {
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
	first, err := repo.FirstDateTransaksi(tx)
	if err != nil {
		panic(err)
	}
	lib.SendJson(map[string]any{"msg": "Success", "first": first}, w)
}
