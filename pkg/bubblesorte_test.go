package day1

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	BubbleSort(arr)
	if !reflect.DeepEqual(arr, []int{3, 4, 7, 9, 42, 69, 420}) {
		t.Errorf("unexpected output")
	}
}
