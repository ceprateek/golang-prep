package linked_list
//
////check if a linked list is a palindrome
//
//func PlayPalindrome() {
//	head := CreatePalindromicLinkedList()
//	printLinkedList(head)
//}
//
//func CreatePalindromicLinkedList() *linkedListNode {
//	head := &linkedListNode{}
//	previous := head
//	for i := 1; i < 4; i++ {
//		node := &linkedListNode{
//			next: nil,
//			val:  i,
//		}
//		previous.next = node
//		previous = node
//	}
//	for i := 3; i >= 1; i-- {
//		node := &linkedListNode{
//			next: nil,
//			val:  i,
//		}
//		previous.next = node
//		previous = node
//	}
//	return head
//}
//
//func findPalindrome(head *linkedListNode) (result bool) {
//	//in a palindrome, first and second half of the lists are reverse of each other
//	//so find middle of the linked list, reverse the second half
//	//compare first and second half
//	//if odd nodes: dont include the mid node
//	//if even nodes: include the 2 mid nodes in each list
//
//	fast := head
//	slow := head
//
//	for fast != nil && fast.next != nil {
//		slow = slow.next
//		fast = fast.next
//	}
//	//fast is nil for even number of elements
//	if fast == nil {
//
//	} else {
//
//	}
//	return result
//}
