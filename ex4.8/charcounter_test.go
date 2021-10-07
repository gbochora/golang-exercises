package charcounter

import (
	"testing"
	"unicode"
)

func TestCountCharsLetters(t *testing.T) {
	bytes := []byte("ააა	ბბabcbb1a2აბგ a")
	got := countChars(bytes, unicode.IsLetter)
	want := map[rune]int{'a': 3, 'b': 3, 'c': 1, 'ა': 4, 'ბ': 3, 'გ': 1}
	if !equal(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func equal(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
