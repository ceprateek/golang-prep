package linked_list

func (l *LinkedList) reverseLinkedList() {
	var previous, runner *LinkedListNode
	runner = l.head
	for runner != nil {
		next := runner.next
		runner.next = previous
		previous = runner
		runner = next
	}
	l.head = previous
}
