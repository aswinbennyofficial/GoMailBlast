package util

import(
	"encoding/csv"
	"os"
	"fmt"
)


type User struct{
	Name string
	Email string
}

func ExtractCsv(filePath string) [][]string {
	// Open the file
	csvfile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	
	for _, record := range records {
		fmt.Println(record)
	}
	return records
}