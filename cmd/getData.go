package cmd

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID         int
	Name       string
	Category   string
	Price      float32
	CreateDate time.Time
}

//функция для сбора данных о продукции, затем полученная информация используется для заполнения csv файла

func GetData(conn *pgx.Conn, connStr string) []Product {
	query := "SELECT * FROM prices"

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error Querying the Table")
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.CreateDate)
		if err != nil {
			log.Printf("Error Fetching Product Details")
		}
		products = append(products, product)

	}

	return products
}
