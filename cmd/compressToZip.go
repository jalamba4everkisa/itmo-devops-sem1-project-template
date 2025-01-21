package cmd

import (
	"archive/zip"
	"io"
	"os"
)

//упаковка csv файла в zip

func CompressToZip(filename string) {
	archive, err := os.Create("data.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w, err := zipWriter.Create("data.csv")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w, f); err != nil {
		panic(err)
	}

	zipWriter.Close()
}
