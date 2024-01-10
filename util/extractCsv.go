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

func ExtractCsv(filePath string) []User {
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
	
	var users []User

	for i, record := range records {
		user := User{
			Name: record[0],
			Email: record[1],
		}
		if i==0{
			continue
		}
		users = append(users, user)
	}
	return users
}