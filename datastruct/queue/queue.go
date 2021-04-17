package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	items []int
	size  int
	sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		items: []int{},
		size:  0,
	}
}

func (q *Queue) Enqueue(t int) {
	q.Lock()
	q.items = append(q.items, t)
	q.size++
	q.Unlock()
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("队列为空")
	}
	q.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.Unlock()
	return item, nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("队列为空")
	}
	q.Lock()
	item := q.items[0]
	q.Unlock()
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
