package day1

type Queue struct {
	Length int
	head   *QueueNode
	tail   *QueueNode
}

type QueueNode struct {
	Value *int
	Next  *QueueNode
}

func (q *Queue) Enqueue(val int) {
	q.Length += 1
	node := &QueueNode{
		Value: &val,
	}

	if q.tail == nil {
		q.head = node
		q.tail = node
	}

	q.tail.Next = node
	q.tail = node
}

func (q *Queue) Deque() *int {
	if q.head == nil {
		return nil
	}
	q.Length -= 1

	head := q.head
	q.head = head.Next

	head.Next = nil

	if q.Length == 0 {
		q.tail = nil
	}

	return head.Value
}

func (q *Queue) Peek() *int {
	if q.head != nil {
		return q.head.Value
	}

	return nil
}
