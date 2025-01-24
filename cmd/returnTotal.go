package cmd

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Total struct {
	TotalItems      int     `json:"total_items"`
	TotalCategories int     `json:"total_categories"`
	TotalPrice      float32 `json:"total_price"`
}

// функция сбора данных после загрузки данных в БД
func ReturnTotal(total_items int, conn *pgxpool.Pool) Total {
	var total Total
	total.TotalItems = total_items
	query := "SELECT COUNT(DISTINCT category) AS total_categories, SUM(price) AS total_price FROM prices;"

	row := conn.QueryRow(context.Background(), query)
	row.Scan(&total.TotalCategories, &total.TotalPrice)

	return total

}
