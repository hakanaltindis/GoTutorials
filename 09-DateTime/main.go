package main

import (
	"fmt"
	"time"
)

func main() {

	// Generate a date and time
	t := time.Date(2016, time.June, 25, 14, 0, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println("------")

	// Get now
	now := time.Now()
	fmt.Printf("The time now is %s\n", now)
	fmt.Println("------")

	// Get day, month, year info from date
	fmt.Println("The day is", t.Day())
	fmt.Println("The month is", t.Month())
	fmt.Println("The year is", t.Year())
	fmt.Println("The day of year is", t.YearDay())

	// add date to time
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println("The tomorrow is", tomorrow)

	// print as formatted
	longFormat := "02-01-2006"
	fmt.Println(tomorrow.Format(longFormat))
}
