package arrays

import (
	"reflect"
	"testing"
)

func TestCopySlice(t *testing.T) {
    got := CopySlice([]int{1, 2, 3})
    want := []int{1, 2, 3}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestAppendToSlice(t *testing.T) {
    got := AppendToSlice([]int{1, 2, 3}, 4, 5, 6)
    want := append([]int{1, 2, 3}, []int{4, 5, 6}...)
    
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }

}

func TestSliceFilter(t *testing.T) {
    got := SliceFilter([]int{1, 2, 3, 4, 5}, func(v int) bool {return v % 2 == 0})
    want := []int{2, 4}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}