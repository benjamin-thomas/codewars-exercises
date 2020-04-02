package words

import (
	"sort"
	"strings"
)

type byLength []string

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// interesting solution
func longestConsecAlt(strarr []string, k int) string {
	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i : i+k], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}
	return largest
}
func LongestConsec(words []string, by int) string {
	size := len(words)
	if by < 1 || size == 0 || by > size {
		return ""
	}

	// Interesting solution, crashes on bad indices though
	return longestConsecAlt(words, by)

	combined := combine(words, by)
	sort.Sort(sort.Reverse(byLength(combined)))

	return combined[0]
}

func combine(words []string, by int) []string {
	badData := by > len(words) || by <= 0
	if badData {
		return []string{}
	}

	newSize := len(words) - by + 1
	arr := make([]string, newSize)

	for i := 0; i < newSize; i++ {
		newWord := ""
		for j := 0; j < by; j++ {
			newWord += words[i+j]
		}
		arr[i] = newWord
	}
	return arr
}