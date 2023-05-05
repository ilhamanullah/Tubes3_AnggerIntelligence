package src

import (
	"fmt"
	"regexp"
	"strconv"
)

func getDay(dateString string) string {

	dateRegex := regexp.MustCompile(`([0-2][0-9]|[1-9]|[3][0-1])(\/)(((0)[1-9])|((1)[0-2])|[1-9])(\/)\d{4}`)
	validate := dateRegex.MatchString(dateString)
	Tanggal := dateRegex.FindAllString(dateString, -1)[0]
	fmt.Println(Tanggal)

	var dayString, monthString, yearString string
	/*
		   kasus 1 : 09/02/2023
		   kasus 2 : 09/2/2023
		   kasus 3 : 9/2/2023
	       kasus 4 : 9/02/2023*/
	if validate {
		fmt.Println("String input valid")
		if Tanggal[2] == '/' {
			if Tanggal[5] == '/' {
				dayString = Tanggal[:2]
				monthString = Tanggal[3:5]
				yearString = Tanggal[6:]
			} else {
				dayString = Tanggal[:2]
				monthString = string(Tanggal[3])
				yearString = Tanggal[5:]
			}
		} else {
			if Tanggal[1] == '/' {
				if Tanggal[4] == '/' {
					dayString = string(Tanggal[0])
					monthString = Tanggal[2:4]
					yearString = Tanggal[5:]
				} else {
					dayString = string(Tanggal[0])
					monthString = string(Tanggal[2])
					yearString = Tanggal[4:]
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
					return "String input tidak valid"
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
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	date = scanner.Text()
// 	fmt.Println(getDay(date))
// }
