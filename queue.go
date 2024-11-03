package circularqueue

type CircularQueue struct {
	values []int
	front  int
	rear   int
	size   int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}
	if q.Empty() {
		q.values[q.rear] = value
		q.size++
		return true
	}
	q.rear = (q.rear + 1) % cap(q.values)
	q.values[q.rear] = value
	q.size++
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.front = (q.front + 1) % cap(q.values)
	q.size--
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
	return q.size == 0
}

func (q *CircularQueue) Full() bool {
	return q.size == cap(q.values)
}
