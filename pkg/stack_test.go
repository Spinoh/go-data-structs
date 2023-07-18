package day1

import "testing"

func TestStack(t *testing.T) {
	list := Stack{}

	list.Push(5)
	list.Push(7)
	list.Push(9)

	if *list.Pop() != 9 {
		t.Errorf("unexpected output")
	}
	if list.Length != 2 {
		t.Errorf("unexpected output")
	}

	list.Push(11)
	if *list.Pop() != 11 {
		t.Errorf("unexpected output")
	}
	if *list.Pop() != 7 {
		t.Errorf("unexpected output")
	}
	if *list.Peek() != 5 {
		t.Errorf("unexpected output")
	}
	if *list.Pop() != 5 {
		t.Errorf("unexpected output")
	}
	if list.Pop() != nil {
		t.Errorf("unexpected output")
	}

	list.Push(69)
	if *list.Peek() != 69 {
		t.Errorf("unexpected output")
	}
	if list.Length != 1 {
		t.Errorf("unexpected output")
	}
}
