package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/callistaenterprise/goblog/quoteservice/model"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {

	quote := model.Quote{
		Text:     "That is a great quote",
		Language: "En",
	}

	data, _ := json.Marshal(quote)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
