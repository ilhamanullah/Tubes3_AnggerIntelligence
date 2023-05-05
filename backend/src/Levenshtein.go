package src

func minimum(a int, b int, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func Levenshtein(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	// fmt.Println(n, m)
	// make matrix with m rows and n columns

	cost := make([][]int, n+1)
	var i, j int
	for i := range cost {
		cost[i] = make([]int, m+1)
	}
	if n == 0 {
		return m
	} else if m == 0 {
		return n
	} else {
		for i = 0; i <= n; i++ {
			cost[i][0] = i
		}
		for j = 0; j <= m; j++ {
			cost[0][j] = j
		}
		for j = 1; j <= m; j++ {
			for i = 1; i <= n; i++ {
				if text[i-1] == pattern[j-1] {
					cost[i][j] = cost[i-1][j-1]
				} else {
					cost[i][j] = minimum(cost[i-1][j]+1, cost[i][j-1]+1, cost[i-1][j-1]+1)
				}
				// print(cost[i][j])
			}
		}
		// fmt.Println(cost[n][m])
		return cost[n][m]
	}
}

func distance(text string, pattern string) float32 {
	dist := Levenshtein(text, pattern)
	// println(dist)
	textLength := len(text)
	patternLength := len(pattern)
	totalLength := textLength + patternLength
	dist = totalLength - dist
	percentage := (float32(dist) / float32(totalLength)) * 100
	return percentage
}

// func main() {
// 	s1 := "siapa presiden indonesia sekarang"
// 	s2 := "siapa presiden indonesia saat ini"
// 	per1 := distance(s1, s2)
// 	fmt.Printf("Persentase kesamaan antara '%s' dan '%s' adalah %.2f%%.\n", s1, s2, per1)
// }
