package main

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func buildLast(pattern string) []int {
	last := make([]int, 128)
	for i := 0; i < 128; i++ {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}

func bmMatch(text string, pattern string) int {
	var last []int = buildLast(pattern)
	n := len(text)
	m := len(pattern)
	i := m - 1

	if i > n-1 {
		return -1
	}

	j := m - 1
	for {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo := last[text[i]]
			i = i + m - min(j, 1+lo)
			j = m - 1
		}

		if i > n-1 {
			break
		}
	}
	return -1
}

func main() {
	text := "ibukota indonesia"
	pattern := "kota indonesia"
	println(bmMatch(text, pattern))
}
