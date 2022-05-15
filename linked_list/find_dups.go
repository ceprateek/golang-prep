package linked_list

import "fmt"

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

type LinkedList struct {
	currentLength int
	head *LinkedListNode
	tail *LinkedListNode
}

func (l *LinkedList) Insert(val int) {
	nodeToInsert := LinkedListNode{
		data: val,
	}
	if l.head == nil {
		l.head = &nodeToInsert
	} else {
		l.tail.next = &nodeToInsert
	}
	l.tail = &nodeToInsert
	l.currentLength++
}

func (l *LinkedList) Display() {
	runner := l.head
	for runner != nil {
		fmt.Println(runner.data)
		runner = runner.next
	}
}

func (l *LinkedList) RemoveDups() {
	cacheOfDups := make(map[int]bool)
	runner := l.head
	var previous *LinkedListNode
	for runner != nil {
		val := runner.data
		if _, ok := cacheOfDups[val]; ok {
			previous.next = runner.next
			runner = runner.next
		} else {
			//add it to the cache
			cacheOfDups[runner.data] = true
			previous = runner
			runner = runner.next
		}
	}
}

func (l *LinkedList) GetKthElementFromLast(k int) int {
	runner := l.head
	for i := 0; i < k; i++ {
		runner = runner.next
	}
	kthPointer := l.head
	for runner != nil {
		runner = runner.next
		kthPointer = kthPointer.next
	}
	return kthPointer.data
}

func (l *LinkedList) partition(p int) {
	runner := l.head
	front := l.head
	end := l.head
	for runner != nil {
		tempNext := runner.next
		if runner.data < p {
			runner.next = front
			front = runner
		} else {
			end.next = runner
			end = runner
		}
		runner = tempNext
	}
	l.head = front
	l.tail = end
}

func PlayLinkedListDupsFinder() {
	l := LinkedList{}   //               1 2 3 6 4 5 8
	l.Insert(6)
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	l.Insert(5)
	l.Insert(8)

	l.RemoveDups()
	l.Display()

	k := l.GetKthElementFromLast(2)
	fmt.Println("--------")
	fmt.Println(k)

	fmt.Println("--------")
	l.partition(4)
	l.Display()
	fmt.Println("--------")
	fmt.Println("--------")
	l.reverseLinkedList()
	l.Display()
}
