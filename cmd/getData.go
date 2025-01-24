package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Product struct {
	ID         int
	Name       string
	Category   string
	Price      float32
	CreateDate time.Time
}

var TotalItems int

//функция для сбора данных о продукции, затем полученная информация используется для заполнения csv файла

func GetData(conn *pgxpool.Pool) []Product {
	query := "SELECT id,name,category,price,create_date FROM prices"

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		fmt.Errorf("error querying prices table: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.CreateDate)
		if err != nil {
			fmt.Errorf("error scanning product row: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("error iterating over rows: %w", err)
	}

	return products
}
