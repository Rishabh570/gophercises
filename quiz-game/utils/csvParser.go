package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func ParseRecords(records [][]string) []Problem {
	problems := make([]Problem, len(records))

	for i, record := range records {
		problems[i] = Problem{
			Question: strings.TrimSpace(record[0]),
			Answer:   strings.TrimSpace(record[1]),
		}
	}

	return problems
}

func ReadCSV(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Failed to open the file: %s\n", filePath)
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	return records
}
