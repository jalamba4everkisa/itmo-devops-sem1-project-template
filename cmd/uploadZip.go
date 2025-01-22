package cmd

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"net/http"
)

//функция обработки zip архива, загруженного через post запрос

func UploadZip(w http.ResponseWriter, r *http.Request) []Product {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body: "+err.Error(), http.StatusBadRequest)
		log.Println("Error reading request body:", err)
	}

	bodyReader := bytes.NewReader(bodyBytes)

	zipReader, err := zip.NewReader(bodyReader, int64(len(bodyBytes)))
	if err != nil {
		http.Error(w, "Error reading zip archive: "+err.Error(), http.StatusBadRequest)
		log.Println("Error reading zip archive:", err)
	}

	var result []Product

	for _, file := range zipReader.File {
		if file.Name == "test_data.csv" {
			result, err = ProcessCSV(file)
			if err != nil {
				http.Error(w, "Error reading CSV inside archive: "+err.Error(), http.StatusBadRequest)
			}
			return result
		}
	}

	http.Error(w, "CSV file not found in zip archive", http.StatusBadRequest)
	log.Println("CSV file not found in zip archive")
	return result
}
