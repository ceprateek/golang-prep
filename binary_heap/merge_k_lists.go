package binary_heap

import (
	"fmt"
)

/*
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	i := 0
	j := len(lists) - 1
	for i < j {
		lists[i] = mergeTwoLists(lists[i], lists[j])
		j--
	}
	return lists[0]
}

//2->4->6->8
//1->2->3->4
func mergeTwoLists(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var result *ListNode
	if a.Val < b.Val {
		result = a
		result.Next = mergeTwoLists(a.Next, b)
	} else {
		result = b
		result.Next = mergeTwoLists(a, b.Next)
	}
	return result
}

func PlayMergeSortedLists() {
	l1 := LinkedList{}
	l1.Insert(1)
	l1.Insert(4)
	l1.Insert(5)

	l2 := LinkedList{}
	l2.Insert(1)
	l2.Insert(3)
	l2.Insert(4)

	l3 := LinkedList{}
	l3.Insert(2)
	l3.Insert(6)

	lists := []*ListNode{l1.head, l2.head, l3.head}
	mergedListHead := mergeKLists(lists)
	for mergedListHead != nil {
		fmt.Println(mergedListHead.Val)
		mergedListHead = mergedListHead.Next
	}
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
	len  int
}

func (l *LinkedList) Insert(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
}
