package reverse

import (
	"unicode/utf8"
)

func reverse(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

func reverseUTF(bytes []byte) {
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		reverse(bytes[i : i+size])
		i += size
	}

	reverse(bytes)
}
