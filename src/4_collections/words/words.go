package words

import (
	"strings"
)

func CountFrequency(words string) map[string]int {
	wordsMap := make(map[string]int, 0)

	for _, word := range strings.Fields(words) {
		trimmedWord := strings.Trim(word, "!?,.-;:\"'`")
		wordsMap[strings.ToLower(trimmedWord)] += 1
	}

	return wordsMap
}