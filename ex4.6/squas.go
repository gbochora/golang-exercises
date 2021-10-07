package squash

import (
	"unicode"
	"unicode/utf8"
)

func squashSpaces(bytes []byte) []byte {
	res := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		curr, size := utf8.DecodeRune(bytes[i:])
		if !unicode.IsSpace(curr) || !unicode.IsSpace(last) {
			res = append(res, bytes[i:i+size]...)
			last = curr
		}
		i += size
	}
	return res
}
