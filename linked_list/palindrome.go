package linked_list

import "fmt"

/*
1-2-2-1 (4/2=2)
1-2-3-2-1 (5/2=2)
*/
func (l *LinkedList) isPalindrome() bool {
	//mid := l.currentLength / 2
	runner := l.head
	fastRunner := l.head
	for fastRunner != nil && fastRunner.next != nil {
		runner = runner.next
		fastRunner = fastRunner.next.next
	}
	last := reverse(runner.next)
	first := l.head
	for last != nil{
		if last.data != first.data{
			return false
		}
		last = last.next
		first = first.next
	}
	return true
}

func reverse(head *LinkedListNode) *LinkedListNode{
	var previous, runner *LinkedListNode
	runner = head
	for runner != nil {
		next := runner.next
		runner.next = previous
		previous = runner
		runner = next
	}
	return previous
}

func PlayPalindromeLL() {
	l := LinkedList{}
	l.Insert(1)
	l.Insert(2)

	result := l.isPalindrome()
	fmt.Println(result)
}
