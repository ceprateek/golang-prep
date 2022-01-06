package linked_list
//
///*
//Input: 1->2->3->4->5->6->7->8->NULL, K = 3
//Output: 3->2->1->6->5->4->8->7->NULL
//Input: 1->2->3->4->5->6->7->8->NULL, K = 5
//Output: 5->4->3->2->1->8->7->6->NULL
//*/
//
//func PlayReverseInGroups() {
//	head := CreateLinkedList()
//	printLinkedList(head)
//	reverseInGroups(3, 9, head)
//	printLinkedList(head)
//}
//
//func reverseInGroups(groupSize, size int, head *linkedListNode) {
//	localHead := head
//	itr := head
//	firstPass := true
//	for i := 1; i <= size; i++ {
//		itr = itr.next
//		if i%groupSize == 0 {
//			next := itr.next
//			itr.next = nil
//			lastNode := reverseLinkedListG(localHead)
//			if firstPass {
//				head = localHead
//				firstPass = false
//			}
//			localHead = next
//			lastNode.next = next
//			itr = lastNode
//		}
//	}
//}
//
////head->1->2->3
////head->3->2->1
//func reverseLinkedListG(head *linkedListNode) *linkedListNode {
//	firstNode := head.next
//	itr := head.next
//	previous := head
//	for itr != nil {
//		next := itr.next
//		itr.next = previous
//		previous = itr
//		itr = next
//	}
//	firstNode.next = nil
//	head.next = previous
//	return firstNode
//}
