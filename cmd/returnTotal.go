package cmd

import (
	"context"
	"fmt"
	"log"
)

type Total struct {
	TotalItems      int     `json:"total_items"`
	TotalCategories int     `json:"total_categories"`
	TotalPrice      float32 `json:"total_price"`
}

// функция сбора данных после загрузки данных в БД
func ReturnTotal() Total {
	u, p, db := LoadEnv()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", u, p, "localhost", "5432", db)
	conn, err := NewPostgres(connStr)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close(context.Background())
	var total Total

	query := "SELECT COUNT(*) AS total_items, COUNT(DISTINCT category) AS total_categories, SUM(price) AS total_price FROM prices;"

	row := conn.QueryRow(context.Background(), query)
	err = row.Scan(&total.TotalItems, &total.TotalCategories, &total.TotalPrice)

	if err != nil {
		fmt.Errorf("query execution or scan failed: %w", err)
	}

	return total

}
