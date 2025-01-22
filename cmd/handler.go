package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//главный хэндлер, роутящий на функции в зависимости от запроса

func PricesHandler(w http.ResponseWriter, r *http.Request) {
	u, p, db := LoadEnv()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", u, p, "localhost", "5432", db)
	conn, err := NewPostgres(connStr)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close(context.Background())
	switch r.Method {
	case http.MethodGet:
		products := GetData(conn, connStr)
		CreateCsv(products)
		CompressToZip("data.csv")
		w.Header().Set("Content-Disposition", "attachment; filename=test-data.zip")
		http.ServeFile(w, r, "data.zip")
		CleanUp()
	case http.MethodPost:
		products := UploadZip(w, r)
		UploadToDb(products, conn, connStr)
		t, err := json.Marshal(ReturnTotal())
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(t)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
