package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func makeError(w http.ResponseWriter, errCode ErrorCode) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	e := Err{Code: errCode}
	responseJSON(w, e)
}

func getURLValue(values url.Values) (string, bool) {
	keys, ok := values["id"]
	if !ok || len(keys[0]) < 1 {
		return "", false
	}
	key := keys[0]
	return key, ok
}

func getURLValueInt(values url.Values) (int, bool) {
	value, ok := getURLValue(values)
	if !ok {
		return 0, false
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}
	return valueInt, true
}
