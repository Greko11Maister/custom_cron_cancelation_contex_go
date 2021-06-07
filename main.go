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
	Ctx          context.Context
	CtxCancel    context.CancelFunc
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx2, cancel2 := context.WithCancel(context.Background())

	companies := []Company{
		Company{Id: "11", BusinessName: "CloudNEt", Duration: time.Minute * 5,
			Ctx:       ctx,
			CtxCancel: cancel,
		},
		Company{Id: "1", BusinessName: "Imaginamos",
			Duration:  time.Minute * 15,
			Ctx:       ctx2,
			CtxCancel: cancel2,
		}}

	// delay := time.Minute

	for _, v := range companies {
		if v.BusinessName == "CloudNEt" {
				go func (c Company) { CancelCronCompany(c) }(v)
		}
		go func(context context.Context, startTime time.Time, delay time.Duration, companyID string, businessName string) {
			for t := range Cron(context, startTime, delay, companyID, businessName) {
				// Perform action here
				log.Println("----------------------------------------------------")
				log.Println("ID: " + t.CompanyID + " BusinessName: " + t.BusinessName)
				log.Println(t.Time.Format("2006-01-02 15:04:05"))
				log.Println("----------------------------------------------------")
			}
		}(v.Ctx, time.Now(), v.Duration, v.Id, v.BusinessName)

	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}

func CancelCronCompany(company Company) {
	time.Sleep(time.Minute * 7)
	log.Println("Cancelando la compañia " + company.BusinessName)
	company.CtxCancel()
}
