package src

// func main() {
// 	expressions := []string{
// 		"(3 + 4) * 5",
// 		"((3 + 4) +* 5) + 6",
// 		"(((3 + 4) + 5) + 6) * 7",
// 		"((a + b) * c) + (d + (e - f))",
// 		"(3 * (4 + 5)",
// 	}

// 	regex := regexp.MustCompile(`^([^\(\)]*|\(([^()]|\(([^()]|\([^()]*\))*\))*\))*$`)
// \s*\d+\s*([-+*\/\^]\s*\d+\s*)*\s*
// 	for _, expr := range expressions {
// 		if regex.MatchString(expr) {
// 			fmt.Printf("%s is a valid expression\n", expr)
// 		} else {
// 			fmt.Printf("%s is not a valid expression\n", expr)
// 		}
// 	}
// }
