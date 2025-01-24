package cmd

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

//функция сборки csv файла

func ProcessCSV(file *zip.File) ([]Product, error) {
	var products []Product

	rc, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("opening file inside zip: %w", err)
	}
	defer rc.Close()

	reader := csv.NewReader(rc)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading csv record: %w", err)
		}
		if len(record) != 5 {
			log.Println("Skipping malformed CSV record (incorrect number of columns):", record)
			continue
		}

		/*id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println("Skipping malformed CSV record (invalid ID):", record, err)
			continue
		}*/

		price, err := strconv.ParseFloat(record[3], 32)
		if err != nil {
			log.Println("Skipping malformed CSV record (invalid Price):", record, err)
			continue
		}

		createDate, err := time.Parse("2006-01-02", record[4])
		if err != nil {
			log.Println("Skipping malformed CSV record (invalid Date):", record, err)
			continue
		}

		product := Product{
			//ID:         id,
			Name:       record[1],
			Category:   record[2],
			Price:      float32(price),
			CreateDate: createDate,
		}

		products = append(products, product)
	}
	return products, nil
}
