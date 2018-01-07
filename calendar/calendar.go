package main

import "fmt"

func calcDayOfWeek(y int, m int, d int) int {
	if m < 3 {
		y--
		m += 12
	}
	return (y + y/4 - y/100 + y/400 + (13*m+8)/5 + d) % 7
}

func isLeapYear(y int, m int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}

func calcDaysInMonth(y int, m int) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	day := days[m-1]
	if isLeapYear(y, m) {
		day += 1
	}
	return day
}

func printCalendar(y int, m int, d int) {
	w := calcDayOfWeek(y, m, 1)
	days := calcDaysInMonth(y, m)

	fmt.Printf("     %d年%2d月\n", y, m)
	fmt.Println("Su Mo Tu We Th Fr Sa")
	for i := 0; i < w; i++ {
		fmt.Print("   ")
	}
	for day := 1; day <= days; day++ {
		fmt.Printf("%2d", day)
		if day%7 == 6 {
			fmt.Println("")
		} else {
			fmt.Print(" ")
		}
		w += 1
	}
}

func main() {
	printCalendar(2018, 1, 8)
}
