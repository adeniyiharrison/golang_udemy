package main

import (
	"fmt"
	"log"

	"github.com/adeniyiharrison/golang_udemy/15_nutritional_details_test/api"
	"github.com/adeniyiharrison/golang_udemy/15_nutritional_details_test/database"
)

func main() {
	upResults := api.ReturnUniversalProducts(1)
	upResults.UpdateProductName()

	fmt.Println(upResults)

	axiomDB := database.ReturnDatabase()
	rows, err := axiomDB.Query(`
		SELECT *
		FROM meijer.products_staging
		LIMIT 3
	`)
	defer rows.Close()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%T\n", rows)

}
