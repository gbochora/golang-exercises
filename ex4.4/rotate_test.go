package rotate

import (
	"reflect"
	"testing"
)

func TestRotateSingleElement(t *testing.T) {
	s := []int{1}
	rotate_ints(s)
	want := []int{1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}

}

func TestRotateSingleEmpty(t *testing.T) {
	s := []int{}
	rotate_ints(s)
	want := []int{}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}

}
func TestRotateEven(t *testing.T) {
	s := []int{1, 2, 3, 4}
	rotate_ints(s)
	want := []int{2, 3, 4, 1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}
}

func TestRotateOdd(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	rotate_ints(s)
	want := []int{1, 2, 3, 4, 0}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}
}
