package linked_list

import (
	"fmt"
	"strconv"
)

type LinkedList1 struct {
	head *node
}

type node struct {
	next *node
	data interface{}
}

func (l *LinkedList1) IsEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *LinkedList1) CreateLinkedList() {
	previous := l.head
	for i := 0; i < 10; i++ {
		node := node{
			data: i,
		}
		if l.IsEmpty() {
			l.head = &node
			previous = &node
		} else {
			previous.next = &node
			previous = &node
		}
	}
}

func (l *LinkedList1) Display() {
	itr := l.head
	var result string
	for itr != nil {
		data := itr.data
		result += " " + strconv.Itoa(data.(int))
		itr = itr.next
	}
	fmt.Println(result)
}

func (l *LinkedList1) GetMiddleLinkedList() interface{} {
	slow := l.head
	fast := l.head
	for fast != nil {
		slow = slow.next
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
	}
	return slow.data
}

//head->b->c->d
//itr->b->c->d

func (l *LinkedList1) Reverse() {
	if l.head == nil || l.head.next == nil{
		return
	}
	previous := l.head
	itr := l.head.next
	for itr != nil {
		nextElement := itr.next
		itr.next = previous
		previous = itr
		itr = nextElement
	}
	l.head.next = nil
	l.head = previous
}
