package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{}
	resp["msg"] = "Hello World"
	jsonData, _ := json.Marshal(resp)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
}
