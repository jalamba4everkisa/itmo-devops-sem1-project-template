package cmd

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// функция для обновления БД новыми данными, полученными из csv файла
func UploadToDb(products []Product, conn *pgxpool.Pool) (TotalItems int) {
	query := "INSERT INTO prices (name,category,price,create_date) VALUES "
	counter := 0
	for i, pr := range products {
		if i == len(products)-1 {
			query += fmt.Sprintf("('%s','%s',%6.1f,'%s')", pr.Name, pr.Category, pr.Price, pr.CreateDate.Format("2006-01-02"))
		} else {
			query += fmt.Sprintf("('%s','%s',%6.1f,'%s'), ", pr.Name, pr.Category, pr.Price, pr.CreateDate.Format("2006-01-02"))
		}
		counter += 1
	}
	err := conn.QueryRow(context.Background(), query)
	fmt.Println(err)
	return counter
}
