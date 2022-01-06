package linked_list
//
//func PlayRemoveLastNthElement()  {
//	head := CreateLinkedList()
//	removeLastNthElement(head, 3)
//	printLinkedList(head)
//}
//
//func removeLastNthElement(head *linkedListNode, N int)  {
//	previous := head
//	itr := head.next
//	var count int
//	for itr != nil{
//		count++
//		if count>N{
//			previous = previous.next
//		}
//		itr = itr.next
//	}
//	if count>N{
//		previous.next = previous.next.next
//	}
//}
