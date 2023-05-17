package main

import (
	"fmt"
	"strings"
)

func findMatchingStrings(n int, strings []string) interface{} {
	matchedIndices := make([]int, 0)
	matchedSet := make(map[string]bool)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if stringsEqualIgnoreCase(strings[i], strings[j]) {
				stringKey := strings[i]
				if _, exists := matchedSet[stringKey]; !exists {
					matchedIndices = append(matchedIndices, i+1, j+1)
					matchedSet[stringKey] = true
				}
			}
		}
	}

	if len(matchedIndices) == 0 {
		return false
	}

	return matchedIndices
}

func stringsEqualIgnoreCase(a, b string) bool {
	return strings.EqualFold(a, b)
}

func main() {
	n1 := 4
	strings1 := []string{"abcd", "acbd", "aaab", "acbd"}
	result1 := findMatchingStrings(n1, strings1)
	fmt.Println("Hasil Test Case :", result1)

	n2 := 5
	strings2 := []string{"pisang", "goreng", "enak", "sekali", "rasanya"}
	result2 := findMatchingStrings(n2, strings2)
	fmt.Println("Hasil Test Case :", result2)
}
