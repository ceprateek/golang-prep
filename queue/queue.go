package queue

import "fmt"

type Queue struct {
	head *node
	tail *node
	size int
}

type node struct {
	next     *node
	previous *node
	data     interface{}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Add(data interface{}) {
	n := &node{}
	n.data = data
	if q.head == nil || q.tail == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		n.previous = q.tail
		q.tail = n
	}
	q.size++
}

func (q *Queue) Remove() interface{} {
	if q.head == nil {
		return nil
	} else if q.head.next == nil {
		data := q.head.data
		q.head = nil
		q.tail = nil
		q.size--
		return data
	} else {
		data := q.head.data
		nextHead := q.head.next
		nextHead.previous = nil
		q.head.next = nil
		q.head = nextHead
		q.size--
		return data
	}
}

func (q *Queue) Display() {
	itr := q.head
	for itr != nil {
		fmt.Printf("%v ", itr.data)
		itr = itr.next
	}
	fmt.Println()
}

func PlayQueue() {
	q := Queue{}
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	fmt.Println(q.Size())
	q.Display()
	q.Remove()
	fmt.Println(q.Size())
	q.Display()

	q.Remove()
	fmt.Println(q.Size())
	q.Display()

	q.Remove()
	fmt.Println(q.Size())
	q.Display()

	q.Remove()
	fmt.Println(q.Size())
	q.Display()

	q.Remove()
	fmt.Println(q.Size())
	q.Display()

}
