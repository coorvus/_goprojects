package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {

	var p Payload

	err := json.NewDecoder(r.Body).Decode(&p)
	log.Println(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if p.Number2 == 0 {
		http.Error(w, "Can not divide by 0", http.StatusBadRequest)
		return
	}

	sum := float64(p.Number1) / float64(p.Number2)

	res := struct {
		Data float64 `json:"data"`
		Msg  string  `json:"msg"`
	}{
		Msg:  "Succeed!",
		Data: math.Round(sum*1000) / 1000,
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
