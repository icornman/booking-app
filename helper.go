package main

import "fmt"

var (
	Info = Color("\033[1;36m%s\033[0m")
	Warn = Color("\033[1;33m%s\033[0m")
	Fata = Color("\033[1;31m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
