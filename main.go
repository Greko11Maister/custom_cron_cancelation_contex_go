package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Company struct {
	Id           string
	BusinessName string
	Duration     time.Duration
}

func main() {
	ctx := context.Background()

	companies := []Company{Company{Id: "11", BusinessName: "CloudNEt", Duration: time.Minute * 5},
		Company{Id: "1", BusinessName: "Imaginamos",
			Duration: time.Minute * 15,
		}}

	// delay := time.Minute

	for _, v := range companies {
		go func(context context.Context, startTime time.Time, delay time.Duration, companyID string, businessName string) {
			for t := range Cron(context, startTime, delay, companyID, businessName) {
				// Perform action here
				log.Println("----------------------------------------------------")
				log.Println("ID: " + t.CompanyID + " BusinessName: " + t.BusinessName)
				log.Println(t.Time.Format("2006-01-02 15:04:05"))
				log.Println("----------------------------------------------------")
			}
		}(ctx, time.Now(), v.Duration, v.Id, v.BusinessName)

	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
