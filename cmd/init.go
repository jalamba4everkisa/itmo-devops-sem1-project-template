package cmd

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgxPool *pgxpool.Pool

//запуск пула, загрузка env

func init() {
	u, p, db := LoadEnv()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", u, p, "localhost", "5432", db)

	PgxPool = NewPostgres(connStr)
}
