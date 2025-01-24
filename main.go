package main

import (
	"net/http"
	"project_sem/cmd"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgxPool *pgxpool.Pool

func main() {

	http.HandleFunc("/api/v0/prices", cmd.PricesHandler)

	http.ListenAndServe(":8080", nil)
}
