package main

import (
	"net/http"
	"project_sem/cmd"
)

func main() {

	http.HandleFunc("/api/v0/prices", cmd.PricesHandler)

	http.ListenAndServe(":8080", nil)
}
