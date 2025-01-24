package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//главный хэндлер, роутящий на функции в зависимости от запроса

func PricesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products := GetData(PgxPool)
		CreateCsv(products)
		CompressToZip("data.csv")
		w.Header().Set("Content-Disposition", "attachment; filename=test-data.zip")
		http.ServeFile(w, r, "data.zip")
		CleanUp()
	case http.MethodPost:
		products := UploadZip(w, r)
		total_items := UploadToDb(products, PgxPool)
		t, err := json.Marshal(ReturnTotal(total_items, PgxPool))
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
