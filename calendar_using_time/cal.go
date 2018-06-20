package main

// timeパッケージを使ったカレンダー

import (
	"fmt"
	"strings"
	"time"

	nowlib "github.com/jinzhu/now"
)

func main() {
	now := time.Now()

	fmt.Printf("%04d-%02d\n", now.Year(), now.Month())

	days := daysIn(now.Year(), now.Month())
	//lastDay := nowlib.EndOfMonth().Day()

	d := nowlib.BeginningOfMonth()
	fmt.Print(strings.Repeat("   ", int(d.Weekday())))
	for ; d.Day() <= days; d = d.AddDate(0, 0, 1) {
		if d.Day() == now.Day() {
			fmt.Printf("\x1b[1m%2d \x1b[0m", d.Day())
		} else if d.Weekday() == time.Saturday {
			fmt.Printf("\x1b[0;34m%2d \x1b[0m", d.Day())
		} else if d.Weekday() == time.Sunday {
			fmt.Printf("\x1b[0;31m%2d \x1b[0m", d.Day())
		} else {
			fmt.Printf("%2d ", d.Day())
		}
		if d.Weekday() == time.Saturday {
			fmt.Print("\n")
		}
		if d.Day() == days {
			break
		}
	}
}

func daysIn(year int, m time.Month) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
