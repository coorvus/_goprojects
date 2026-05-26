package main

import (
	"coorvus/calculator/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /add", handlers.SumHandler)
	http.HandleFunc("POST /substract", handlers.SubstractHandler)
	http.HandleFunc("POST /multiply", handlers.MultiplyHandler)
	http.HandleFunc("POST /divide", handlers.DivideHandler)

	s := &http.Server{
		Addr: ":8000",
	}

	log.Println("Server start at port 8000")
	s.ListenAndServe()
}
