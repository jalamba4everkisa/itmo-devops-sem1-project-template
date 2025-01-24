package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

//функция для коннекта к базе постгри

func NewPostgres(connStr string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(dbpool.Ping(context.Background()))

	return dbpool
}
