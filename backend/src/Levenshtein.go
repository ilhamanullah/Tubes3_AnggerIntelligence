package main

import (
	"fmt"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func distance(text string, pattern string) float32 {
	dist := levenshtein.DistanceForStrings([]rune(text), []rune(pattern), levenshtein.DefaultOptions)
	// println(dist)
	textLength := len(text)
	patternLength := len(pattern)
	totalLength := textLength + patternLength
	dist = totalLength - dist
	percentage := (float32(dist) / float32(totalLength)) * 100
	return percentage
}

func main() {
	s1 := "siapa presiden indonesia sekarang"
	s2 := "siapa presiden indonesia saat ini"
	per1 := distance(s1, s2)
	fmt.Printf("Persentase kesamaan antara '%s' dan '%s' adalah %.2f%%.\n", s1, s2, per1)
}
