package squash

import (
	"reflect"
	"testing"
)

func TestSquashEmpty(t *testing.T) {
	b := []byte("")
	got := squashSpaces(b)
	want := []byte("")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSquashEmptyOnlySpaces(t *testing.T) {
	b := []byte("    ")
	got := squashSpaces(b)
	want := []byte(" ")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSquashSpacesAtStart(t *testing.T) {
	b := []byte("    ABC")
	got := squashSpaces(b)
	want := []byte(" ABC")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSquashSpaces(t *testing.T) {
	b := []byte("    A  B C  ")
	got := squashSpaces(b)
	want := []byte(" A B C ")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSquashSpacesUnicode(t *testing.T) {
	b := []byte("   ა  A   ბ  B C  ")
	got := squashSpaces(b)
	want := []byte(" ა A ბ B C ")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
