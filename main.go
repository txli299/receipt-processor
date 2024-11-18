package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/txli299/receipt-processor/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
