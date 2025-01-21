package cmd

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// функция для обновления БД новыми данными, полученными из csv файла
func UploadToDb(products []Product, conn *pgx.Conn, connStr string) {
	query := "INSERT INTO prices (id,name,category,price,create_date) VALUES "

	for i, pr := range products {
		if i == len(products)-1 {
			query += fmt.Sprintf("(%x,'%s','%s',%6.1f,'%s')", pr.ID, pr.Name, pr.Category, pr.Price, pr.CreateDate.Format("2006-01-02"))
		} else {
			query += fmt.Sprintf("(%x,'%s','%s',%6.1f,'%s'), ", pr.ID, pr.Name, pr.Category, pr.Price, pr.CreateDate.Format("2006-01-02"))
		}
	}
	err := conn.QueryRow(context.Background(), query)
	fmt.Println(err)
}
