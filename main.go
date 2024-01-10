package main

import (
	
	"log"
	"github.com/aswinbennyofficial/GoMailBlast/util"
)

func main(){
	// Extracting CSV
	records := util.ExtractCsv("./data/data.csv")
	log.Println(records)

	// Sending email
	util.SendBulkEmail(records)
	
	
}