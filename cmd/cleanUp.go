package cmd

import (
	"log"
	"os"
)

// Подчистка мусора, после создания и упаковки csv в zip
func CleanUp() {
	os.Remove("data.csv")
	os.Remove("data.zip")

}

func Remove(filename string) {
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}
