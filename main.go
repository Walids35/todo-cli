package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const filename string = "data.csv"

func main(){
	LoadTasks(filename)
}

func LoadTasks(filename string) {
	data, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	reader := csv.NewReader(data)
	records, err := reader.ReadAll()

	if err != nil{
		fmt.Println("Error reading records")
	}

	for _,record := range records{
		fmt.Println(record)
	}
}
