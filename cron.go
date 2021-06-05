package main

import (
	"context"
	"time"
)

func Cron(ctx context.Context, startTime time.Time, delay time.Duration, companyID string, businessName string) <-chan ConfigStructure {
	// Create the channel which we will return
	stream := make(chan ConfigStructure, 1)

	// Calculating the first start time in the future
	// Need to check if the time is zero (e.g. if time.Time{} was used)
	if !startTime.IsZero() {
		diff := time.Until(startTime)
		if diff < 0 {
			total := diff - delay
			times := total / delay * -1

			startTime = startTime.Add(times * delay)
		}
	}

	// Run this in a goroutine, or our function will block until the first event
	go func() {

		// Run the first event after it gets to the start time
		t := <-time.After(time.Until(startTime))
		stream <- ConfigStructure{
			Time:         t,
			CompanyID:    companyID,
			BusinessName: businessName,
		}

		// Open a new ticker
		ticker := time.NewTicker(delay)
		// Make sure to stop the ticker when we're done
		defer ticker.Stop()

		// Listen on both the ticker and the context done channel to know when to stop
		for {
			select {
			case t2 := <-ticker.C:
				stream <- ConfigStructure{
					Time:         t2,
					CompanyID:    companyID,
					BusinessName: businessName,
				}
			case <-ctx.Done():
				close(stream)
				return
			}
		}
	}()

	return stream
}
