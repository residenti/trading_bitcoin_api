package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/residenti/trading_bitcoin_api/app/models"
	"github.com/residenti/trading_bitcoin_api/config"
)

func StartServer() error {
	http.HandleFunc("/candle/", apiMakeHandler(apiCandleHandler))
	return http.ListenAndServe(fmt.Sprintf(":%d", config.List.Port), nil)
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/candle/$")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found.", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func apiCandleHandler(w http.ResponseWriter, r *http.Request) {
	productCode := r.URL.Query().Get("product_code")
	if productCode == "" {
		APIError(w, "No product_code param.", http.StatusBadRequest)
		return
	}

	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 1000 {
		limit = 10000
	}

	duration := r.URL.Query().Get("duration")
	if duration == "" {
		duration = "1s" // TODO 設定ファイルを読み取る.
	}
	durationTime := config.List.Durations[duration]

	dfCandle, err := models.GetDataFrameCandle(productCode, durationTime, limit)
	if err != nil {
		log.Fatalln(err)
	}

	js, err := json.Marshal(dfCandle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
