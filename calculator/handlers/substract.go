package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func SubstractHandler(w http.ResponseWriter, r *http.Request) {

	var p Payload

	err := json.NewDecoder(r.Body).Decode(&p)
	log.Println(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sum := p.Number1 - p.Number2

	res := struct {
		Data int    `json:"data"`
		Msg  string `json:"msg"`
	}{
		Msg:  "Succeed!",
		Data: sum,
	}

	w.Header().Set("Content-Type", "application/json")
	bt, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Failed to add two number", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bt)

}
