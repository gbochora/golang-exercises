package duplicats

import (
	"reflect"
	"testing"
)

func TestEliminateEdjacentDupEmpty(t *testing.T) {
	strings := []string{""}
	want := []string{""}
	got := eliminateAdjacentDup(strings)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", strings, want)
	}
}

func TestEliminateEdjacentDup1(t *testing.T) {
	strings := []string{"a", "a", "a", "b"}
	want := []string{"a", "b"}
	got := eliminateAdjacentDup(strings)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEliminateEdjacentDup2(t *testing.T) {
	strings := []string{"a", "a", "a"}
	want := []string{"a"}
	got := eliminateAdjacentDup(strings)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEliminateEdjacentDup3(t *testing.T) {
	strings := []string{"a", "a", "a", "b", "b"}
	want := []string{"a", "b"}
	got := eliminateAdjacentDup(strings)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
