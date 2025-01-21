package cmd

import (
	"fmt"
	"os"
)

const (
	YYYYMMDD = "2006-01-02"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//функция для создания csv файла и заполнения его построчно

func CreateCsv(products []Product) {

	f, e1 := os.Create("data.csv")
	check(e1)
	defer f.Close()
	_, e2 := f.WriteString("id,name,category,price,create_date\n")
	check(e2)

	for _, pr := range products {
		line := fmt.Sprintf("%x,%s,%s,%6.2f,%s\n", pr.ID, pr.Name, pr.Category, pr.Price, pr.CreateDate.Format(YYYYMMDD))
		_, e3 := f.WriteString(line)
		check(e3)
	}
}
