//go:build ignore

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){

	
	file, err := os.Open("data/newfile.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}


	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := strings.ReplaceAll(scanner.Text(), ",", "")
	text = strings.ReplaceAll(text, ".", "")
	inputs := strings.Split(text, " ")

	sortedWords := SortWord(inputs)
	matchWordsLine, otherWords := searchCSV(lines, sortedWords)

	fmt.Println(otherWords)

	groupByCefr := make(map[string][][]string)

	for _, line := range matchWordsLine{
		field := line[2]
		groupByCefr[field] = append(groupByCefr[field], line[:2])
	}

	for _, field := range []string{"A1", "A2", "B1", "B2", "C1", "C2"}{
		lines, ok :=groupByCefr[field]
		if ok {
			fmt.Printf("\nField %s:\n", field)
			for _, line := range lines {
				fmt.Printf("%s %s\n", line[0], line[1])
			}
		}
	}
}
func bindCSV(){
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

func SortWord(words []string) []string{
	var sortedWords []string
	seen := make(map[string]bool)
	for _, word := range words{
		if !seen[word]{
			sortedWords = append(sortedWords, word)
			seen[word]= true
		}
	}
	return sortedWords
}

func searchCSV(lines [][]string, words []string)([][]string, []string){
	var matchWordsLine [][]string
	var otherWords []string
	for _, word := range words{
		flag := false
		for _, line := range lines{
			if word ==  line[0]{
				matchWordsLine = append(matchWordsLine, line)
				flag = true
			} 
		}
		if !flag{
				otherWords = append(otherWords, word)
			}
	}
	return matchWordsLine, otherWords
}

