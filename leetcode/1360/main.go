package main

import (
	"fmt"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	time1, _ := time.Parse("2006-01-02", date1)
	time2, _ := time.Parse("2006-01-02", date2)

	diffHours := time2.Sub(time1).Hours()

	return int(diffHours / 24)
}

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
}
