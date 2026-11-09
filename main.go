package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	file, err := excelize.OpenFile("RentPerSquareFootAug2026.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	all_rows, err := file.GetRows("RentPerSquareFootAug2026")
	if err != nil {
		fmt.Println("Error getting Rows", err)
	}
	for _, row := range all_rows {
		//the titles all have one real cell, but the data we want all has many cells, so only do work
		//when there is more than one cell in the row
		if len(row) > 1 {
			fmt.Println(row[0], " \t:\t ", row[2])
		}
	}
}
