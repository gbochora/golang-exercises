package reverse

import (
	"reflect"
	"testing"
)

func TestReverseCharsEmpty(t *testing.T) {
	got := []byte("")
	reverseUTF(got)
	want := []byte("")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseChars(t *testing.T) {
	got := []byte("აaბbგc")
	reverseUTF(got)
	want := []byte("cგbბaა")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
