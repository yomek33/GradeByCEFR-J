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
	data = append(data, readCSV("cefrj-vocabulary-profile-1.5.csv")...)
	data = append(data, readCSV("octanove-vocabulary-profile-c1c2-1.0.csv")...)

	newFile, err := os.Create("newfile.csv")
	if err != nil {
        log.Fatal(err)
    }
	writer := csv.NewWriter(newFile)
    defer writer.Flush()

	for _, record := range data {
		if record[2] != "CEFR"{
			err := writer.Write(record[:3])
			if err != nil {
				panic(err)
			}
		}
    }
	fmt.Println("bind!!")
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