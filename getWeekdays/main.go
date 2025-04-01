package main

import (
	"fmt"
	"time"
)

func getWeekDays() {
	t := time.Now()
	
	switch {
	case t.Weekday() == time.Sunday:
		fmt.Println("Today is", t.Weekday())
	case t.Weekday() == time.Monday:
		fmt.Println("Today is", t.Weekday())
	case t.Weekday() == time.Tuesday:
		fmt.Println("Today is", t.Weekday())
	case t.Weekday() == time.Wednesday:
		fmt.Println("Today is", t.Weekday())
	case t.Weekday() == time.Thursday:
		fmt.Println("Today is", t.Weekday())
	case t.Weekday() == time.Friday:
		fmt.Println("Today is", t.Weekday())
	case t.Weekday() == time.Saturday:
		fmt.Println("Today is", t.Weekday())	
	}
}

func main () {
	getWeekDays()
}