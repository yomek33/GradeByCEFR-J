//go:build ignore

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main(){
	data := make([][]string, 0)
	data = append(data, readCSV("data/cefrj-vocabulary-profile-1.5.csv")...)
	data = append(data, readCSV("data/cefrj-vocabulary-profile-1.5.csv")...)

	newFile, err := os.Create("data/newfile.csv")
	if err != nil {
        log.Fatal(err)
    }
	writer := csv.NewWriter(newFile)
    defer writer.Flush()

	for _, record := range data {
        err := writer.Write(record)
		if err != nil {
			panic(err)
		}
    }
	fmt.Println("/nbind!!")
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	reader :=csv.NewReader(file)
	records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

	return records
}