package day1

import (
	"testing"
)

func TestQueue(t *testing.T) {
	list := Queue{}

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	popped := *list.Deque()
	if popped != 5 {
		t.Errorf("unexpected output")
	}

	if list.Length != 2 {
		t.Errorf("unexpected length")
	}

	list.Enqueue(11)

	if *list.Deque() != 7 {
		t.Errorf("unexpected output")
	}
	if *list.Deque() != 9 {
		t.Errorf("unexpected output")
	}
	if *list.Peek() != 11 {
		t.Errorf("unexpected output")
	}
	if *list.Deque() != 11 {
		t.Errorf("unexpected output")
	}
	if list.Deque() != nil {
		t.Errorf("unexpected output")
	}

	list.Enqueue(69)
	if *list.Peek() != 69 {
		t.Errorf("unexpected output")
	}
	if list.Length != 1 {
		t.Errorf("unexpected output")
	}
}
