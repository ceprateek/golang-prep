package linked_list
//
///*
//Input: 10->20->30->40->50->60
//k=4
//Result: 50->60->10->20->30->40
//*/
//
//func PlayRotate() {
//
//}
//
//func rotate(head *linkedListNode, k int) {
//	if k == 0 {
//		return
//	}
//	current := head
//	count := 1
//	for count < k && current != nil {
//		current = current.next
//		count++
//	}
//	if current == nil {
//		return
//	}
//	kthNode := current
//
//	for current.next != nil {
//		current = current.next
//	}
//	current.next = head
//	head = kthNode.next
//	kthNode.next = nil
//}
//
//// CreateLinkedListHeadPointer here the head is not a new element but simply a pointer to the first element
//// head.next would point to the 2nd element and head to first
//func CreateLinkedListHeadPointer() *linkedListNode {
//	var head *linkedListNode
//	for i := 10; i > 0; i-- {
//		node := &linkedListNode{
//			next: nil,
//			val:  i,
//		}
//		node.next = head
//		head = node
//	}
//	return head
//}
