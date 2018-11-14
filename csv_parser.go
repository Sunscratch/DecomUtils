package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readFromFile(file string, separator rune) [][]string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Promblem opening file: "+file, err)
	}
	r := csv.NewReader(strings.NewReader(string(data)))
	r.Comma = separator
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Promblem reading "+file+" file: ", err)
	}
	return records
}

func writeToFile(csvData [][]string, filename string, delimiter rune) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Could not create file "+filename, err)
	}
	w := csv.NewWriter(file)
	w.Comma = delimiter
	w.WriteAll(csvData)
	if err := w.Error(); err != nil {
		log.Fatal("Could not write data "+filename, err)
	}
}
