package main

import (
	"fmt"
)

/*

You are given the following information, but you may prefer to do some research for yourself.

* 1 Jan 1900 was a Monday.
* Thirty days has September,
  April, June and November.
  All the rest have thirty-one,
  Saving February alone,
  Which has twenty-eight, rain or shine.
  And on leap years, twenty-nine.
* A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

https://projecteuler.net/problem=19

*/

func main() {
	months := [13]int{0,31,28,31,30,31,30,31,31,30,31,30,31}

	weekday := 0
	day := 1
	month := 1
	year := 1900

	count := 0

	for i := 0; ; i++ {
		if day == months[month] {
			day = 1
			month = changeMonth(month)
			if month == 1 {
				year++
				if year % 4 == 0 && ( year % 100 != 0 || year % 400 == 0) {
					months[2] = 29
					fmt.Println(year)
				} else {
					months[2] = 28
				}
			}

			if year > 1900 && day == 1 && weekday == 6 {
				count++
			}
		}

		day++
		weekday = changeDay(weekday)

		if( year == 2000 && month == 12 && day == 31 ) {
			break
		}
	}

	//fmt.Printf("%d.%d.%d weekday : %d\n", day, month, year, weekday)
	fmt.Println(count)
}

func changeDay(weekday int) int {
	if weekday == 6 {
		return 0
	}
	return weekday + 1
}

func changeMonth(month int) int {
	if month == 12 {
		return 1
	}
	return month + 1
}
