package algors

import "container/list"

type Queue struct {
	list list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: *list.New(),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	e := q.list.Front()
	q.list.Remove(e)
	return e
}

func (q *Queue) Enqueue(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	e := q.list.Back()
	q.list.Remove(e)
	return e
}
