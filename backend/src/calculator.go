package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 25+8-4*4/2^2

func countOperator(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' || input[i] == '/' || input[i] == '^' {
			count++
		}
	}
	return count
}

func isOperator(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' || input[i] == '/' || input[i] == '^' {
			if input[0] == '-' && countOperator(input) == 1 {
				return false
			}
			return true
		}
	}
	return false
}

func isKuadrat(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '^' {
			return true
		}
	}
	return false
}

func isKali(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '*' {
			return true
		}
	}
	return false
}

func isBagi(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '/' {
			return true
		}
	}
	return false
}

func isKaliBagi(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '*' || input[i] == '/' {
			return true
		}
	}
	return false
}

func isTambah(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '+' {
			return true
		}
	}
	return false
}

func isKurang(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == '-' {
			return true
		}
	}
	return false
}

func getVal(calculator string) float64 {

	calRegex := regexp.MustCompile(`^\s*\d+\s*([-+*\/\^]\s*\d+\s*)*\s*$`)
	validate := calRegex.MatchString(calculator)
	var i int
	if validate {
		isSpace := false
		// var count, sum int
		// return 0
		for i = 0; i < len(calculator); i++ {
			if calculator[i] == ' ' {
				isSpace = true
				break
			}
		}
		if isSpace {
			calculator = strings.ReplaceAll(calculator, " ", "")
		}

		for isOperator(calculator) {
			var count int
			for i = 0; i < len(calculator); i++ {
				if calculator[i] == ' ' {
					isSpace = true
				}
				if calculator[i] == '+' || calculator[i] == '-' || calculator[i] == '*' || calculator[i] == '/' || calculator[i] == '^' {
					count++
				}
			}
			/* keterengan:
			+ -> 1
			- -> 2
			* -> 3
			/ -> 4
			^ -> 5*/
			operator := make([][]int, count)
			for j := range operator {
				operator[j] = make([]int, 3)
			}
			k := 0
			for i = 0; i < len(calculator); i++ {
				if calculator[i] == '+' {
					operator[k][0] = 1
					operator[k][1] = i
					operator[k][2] = 1
					k++
				} else if calculator[i] == '-' {
					operator[k][0] = 2
					operator[k][1] = i
					operator[k][2] = 1
					k++
				} else if calculator[i] == '*' {
					operator[k][0] = 3
					operator[k][1] = i
					operator[k][2] = 1
					k++
				} else if calculator[i] == '/' {
					operator[k][0] = 4
					operator[k][1] = i
					operator[k][2] = 1
					k++
				} else if calculator[i] == '^' {
					operator[k][0] = 5
					operator[k][1] = i
					operator[k][2] = 1
					k++
				}
			}

			if isKuadrat(calculator) {
				for i = 0; i < len(operator); i++ {
					if operator[i][0] == 5 {
						if operator[i][2] == 1 {
							var num1, num2 float64
							var str1, str2 string
							if len(operator) == 1 {
								str1 = calculator[0:operator[i][1]]
								// fmt.Printf("str1: %s\n", str1)
								num1, _ = strconv.ParseFloat(str1, 8)
								// fmt.Println("num1 - 1 kuad:", str1)
								str2 = calculator[operator[i][1]+1:]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num2 - 1 kuad:", str2)
							} else if i == 0 {
								str1 = calculator[0:operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 2 kuad:", num1)
								// fmt.Println("num2 - 2 kuad:", num2)

							} else if i == len(operator)-1 {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : len(calculator)]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 3 kuad:", num1)
								// fmt.Println("num2 - 3 kuad:", num2)
							} else {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 4 kuad:", num1)
								// fmt.Println("num2 - 4 kuad:", num2)
							}
							sum := math.Pow(float64(num1), float64(num2))
							calculator = strings.Replace(calculator, str1+"^"+str2, strconv.Itoa(int(sum)), 1)
							break
						}
					}
				}
				continue
			} else if isKaliBagi(calculator) {
				for i = 0; i < len(operator); i++ {
					if operator[i][0] == 3 {
						if operator[i][2] == 1 {
							var num1, num2 float64
							var str1, str2 string
							if len(operator) == 1 {
								str1 = calculator[0:operator[i][1]]
								// fmt.Printf("str1: %s\n", str1)
								num1, _ = strconv.ParseFloat(str1, 8)
								// fmt.Println("num1 - 1 kali:", str1)
								str2 = calculator[operator[i][1]+1:]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num2 - 1 kali:", str2)
							} else if i == 0 {
								str1 = calculator[0:operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 2 kali:", num1)
								// fmt.Println("num2 - 2 kali:", num2)

							} else if i == len(operator)-1 {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : len(calculator)]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 3 kali:", num1)
								// fmt.Println("num2 - 3 kali:", num2)
							} else {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 4 kali:", num1)
								// fmt.Println("num2 - 4 kali:", num2)
							}
							sum := num1 * num2
							calculator = strings.Replace(calculator, str1+"*"+str2, fmt.Sprintf("%f", sum), 1)
							// fmt.Println("calculator kali:", calculator)
							break
						}
					} else if operator[i][0] == 4 {
						if operator[i][2] == 1 {
							var num1, num2 float64
							var str1, str2 string
							if len(operator) == 1 {
								str1 = calculator[0:operator[i][1]]
								// fmt.Printf("str1: %s\n", str1)
								num1, _ = strconv.ParseFloat(str1, 8)
								// fmt.Println("num1 - 1 bagi:", str1)
								str2 = calculator[operator[i][1]+1:]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num2 - 1 bagi:", str2)
							} else if i == 0 {
								str1 = calculator[0:operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 2 bagi:", num1)
								// fmt.Println("num2 - 2 bagi:", num2)

							} else if i == len(operator)-1 {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								// fmt.Println(operator[i][1] + 1)
								// fmt.Println(len(calculator) - 1)
								str2 = calculator[operator[i][1]+1 : len(calculator)]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 3 bagi:", num1)
								// fmt.Println("num2 - 3 bagi:", num2)
							} else {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 4 bagi:", num1)
								// fmt.Println("num2 - 4 bagi:", num2)
							}
							sum := num1 / num2
							calculator = strings.Replace(calculator, str1+"/"+str2, fmt.Sprintf("%f", sum), 1)
							// fmt.Println("calculator bagi: ", calculator)
							break
						}
					}
				}
				continue
			} else if isTambah(calculator) {
				for i = 0; i < len(operator); i++ {
					if operator[i][0] == 1 {
						if operator[i][2] == 1 {
							var num1, num2 float64
							var str1, str2 string
							if len(operator) == 1 {
								str1 = calculator[0:operator[i][1]]
								// fmt.Printf("str1: %s\n", str1)
								num1, _ = strconv.ParseFloat(str1, 8)
								// fmt.Println("num1 - 1 tambah:", str1)
								str2 = calculator[operator[i][1]+1:]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num2 - 1 tambah:", str2)
							} else if i == 0 {
								str1 = calculator[0:operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 2 tambah:", num1)
								// fmt.Println("num2 - 2 tambah:", num2)

							} else if i == len(operator)-1 {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : len(calculator)]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 3 tambah:", num1)
								// fmt.Println("num2 - 3 tambah:", num2)
							} else {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 4 tambah:", num1)
								// fmt.Println("num2 - 4 tambah:", num2)
							}
							sum := num1 + num2
							// fmt.Println("sum:", sum)
							calculator = strings.Replace(calculator, str1+"+"+str2, fmt.Sprintf("%f", sum), 1)
							// fmt.Println("calculator tambah:", calculator)
							break
						}
					}
				}
				continue
			} else if isKurang(calculator) {
				for i = 0; i < len(operator); i++ {
					if operator[i][0] == 2 {
						if operator[i][2] == 1 {
							var num1, num2 float64
							var str1, str2 string
							if len(operator) == 1 {
								str1 = calculator[0:operator[i][1]]
								// fmt.Printf("str1: %s\n", str1)
								num1, _ = strconv.ParseFloat(str1, 8)
								// fmt.Println("num1 - 1 kurang:", str1)
								str2 = calculator[operator[i][1]+1:]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num2 - 1 kurang:", str2)
							} else if i == 0 {
								str1 = calculator[0:operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 2 kurang:", num1)
								// fmt.Println("num2 - 2 kurang:", num2)

							} else if i == len(operator)-1 {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : len(calculator)]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 3 kurang:", num1)
								// fmt.Println("num2 - 3 kurang:", num2)
							} else {
								str1 = calculator[operator[i-1][1]+1 : operator[i][1]]
								num1, _ = strconv.ParseFloat(str1, 8)
								str2 = calculator[operator[i][1]+1 : operator[i+1][1]]
								num2, _ = strconv.ParseFloat(str2, 8)
								// fmt.Println("num1 - 4 kurang:", num1)
								// fmt.Println("num2 - 4 kurang:", num2)
							}
							sum := num1 - num2
							calculator = strings.Replace(calculator, str1+"-"+str2, fmt.Sprintf("%f", sum), 1)
							// fmt.Println("calculator kurang:", calculator)
							break
						}
					}
				}
				continue
			}

			if isOperator(calculator) == false {
				break
			}

		}

	}

	// fmt.Println(calculator)
	hasil, _ := strconv.ParseFloat(calculator, 8)
	return hasil
}

func main() {
	var calculator string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	calculator = scanner.Text()
	fmt.Println(getVal(calculator))
}
