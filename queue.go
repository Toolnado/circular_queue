package circularqueue

type CircularQueue struct {
	values []int
	front  int
	rear   int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		front:  -1,
		rear:   -1,
	}
}

func (q *CircularQueue) increaseFront() {
	q.front = (q.front + 1) % cap(q.values)
}

func (q *CircularQueue) increaseRear() {
	q.rear = (q.rear + 1) % cap(q.values)
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	if q.front == -1 {
		q.increaseFront()
	}
	q.increaseRear()
	q.values[q.rear] = value
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.values[q.front] = 0
	if q.front == q.rear {
		q.rear = -1
		q.front = -1
	} else {
		q.increaseFront()
	}
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

func (q *CircularQueue) Empty() bool {
	return q.front == -1
}

func (q *CircularQueue) Full() bool {
	return q.front == q.rear+1 ||
		(q.front == 0 && q.rear == cap(q.values)-1)
}
