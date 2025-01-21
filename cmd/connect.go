package cmd

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//функция для коннекта к базе постгри

func NewPostgres(connStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
