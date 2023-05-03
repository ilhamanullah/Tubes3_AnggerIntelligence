package main

func kmpMatch(text string, pattern string) int {
	n := len(text)
	m := len(pattern)

	b := computeBorder(pattern)
	i := 0
	j := 0

	for i < n {
		if text[i] == pattern[j] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else {
			if j > 0 {
				j = b[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func computeBorder(pattern string) []int {
	b := make([]int, len(pattern))
	m := len(pattern)
	i := 1
	j := 0

	for i < m {
		if pattern[i] == pattern[j] {
			b[i] = j + 1
			i++
			j++
		} else {
			if j > 0 {
				j = b[j-1]
			} else {
				b[i] = 0
				i++
			}
		}
	}
	return b
}

func main() {
	text := "apakah anda suka"
	pattern := "anda suka"
	println(kmpMatch(text, pattern))
}
