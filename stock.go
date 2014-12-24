package main

import (
	"encoding/csv"
	//log "github.com/cihub/seelog"
	"net/http"
)

func createquery(symbol string) string {
	s := "http://download.finance.yahoo.com/d/quotes.csv?s=" + symbol
	return s + "&f=nsl1op"
}

func runquery(query string) map[string]string {
	res, _ := http.Get(query)
	reader := csv.NewReader(res.Body)
	rawvalues, _ := reader.Read()

	values := make(map[string]string)
	values["name"] = rawvalues[0]
	values["symbol"] = rawvalues[1]
	values["latest"] = rawvalues[2]
	values["open"] = rawvalues[3]
	values["close"] = rawvalues[4]

	return values
}

func makestring(values map[string]string) string {
	s := ""
	return s + values["name"] + " [" + values["symbol"] +
		"] latest: " + values["latest"] + " open: " +
		values["open"] + " close: " + values["close"]
}

func stock(symbol string) string {
	return makestring(runquery(createquery(symbol)))
}
