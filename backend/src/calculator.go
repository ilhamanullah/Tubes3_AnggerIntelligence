package src

import (
	"regexp"
)

func getVal(calculator string) string {

	calRegex := regexp.MustCompile(`^\s*\d+\s*([-+*\/\^]\s*\d+\s*)*\s*$`)
	validate := calRegex.MatchString(calculator)
	if validate {
		return "String input valid"
	}
	return "String input tidak valid"
}

// func main() {
// 	var calculator string
// 	fmt.Scan(&calculator)
// 	fmt.Println(getVal(calculator))
// }
