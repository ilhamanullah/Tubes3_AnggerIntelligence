package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getDay(dateString string) string {

	dateRegex := regexp.MustCompile(`^([0-2][0-9]|[1-9]|[3][0-1])(\/)(((0)[1-9])|((1)[0-2])|[1-9])(\/)\d{4}$`)
	validate := dateRegex.MatchString(dateString)
	var dayString, monthString, yearString string
	/*
		   kasus 1 : 09/02/2023
		   kasus 2 : 09/2/2023
		   kasus 3 : 9/2/2023
	       kasus 4 : 9/02/2023*/
	if validate {
		fmt.Println("String input valid")
		if dateString[2] == '/' {
			if dateString[5] == '/' {
				dayString = dateString[:2]
				monthString = dateString[3:5]
				yearString = dateString[6:]
			} else {
				dayString = dateString[:2]
				monthString = string(dateString[3])
				yearString = dateString[5:]
			}
		} else {
			if dateString[1] == '/' {
				if dateString[4] == '/' {
					dayString = string(dateString[0])
					monthString = dateString[2:4]
					yearString = dateString[5:]
				} else {
					dayString = string(dateString[0])
					monthString = string(dateString[2])
					yearString = dateString[4:]
				}
			}
		}
		hari, _ := strconv.Atoi(dayString)
		bulan, _ := strconv.Atoi(monthString)
		tahun, _ := strconv.Atoi(yearString)

		if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
			if hari == 31 {
				return "String input tidak valid"
			}
		}
		if bulan == 2 {
			if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
				if hari > 29 {
					return "String input tidak valid 1"
				}
			} else {
				if hari > 28 {
					return "String input tidak valid"
				}
			}
		}

		a := (14 - bulan) / 12
		y := tahun - a
		m := bulan + 12*a - 2
		d := (hari + y + y/4 - y/100 + y/400 + (31*m)/12) % 7

		days := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
		return "Tanggal tersebut berada di hari " + days[d]

	} else {
		return "String input tidak valid"
	}
}

// func main() {
// 	//make string input
// 	var date string
// 	fmt.Scan(&date)
// 	fmt.Println(getDay(date))
// 	// fmt.Println(getDay("29/2/2024"))
// }
