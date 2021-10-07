package charcounter

import "unicode/utf8"

func countChars(bytes []byte, isCharInCategory func(rune) bool) map[rune]int {
	counts := make(map[rune]int)
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		if isCharInCategory(r) {
			counts[r]++
		}
		i += size
	}
	return counts
}
