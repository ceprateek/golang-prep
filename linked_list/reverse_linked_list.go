package linked_list

//import "fmt"
//
//type linkedListNode struct {
//	next *linkedListNode
//	val  int
//}
//
//func CreateLinkedList() *linkedListNode{
//	head := &linkedListNode{}
//	previous := head
//	for i := 1; i < 10; i++ {
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
//func PlayLinkedListReverse() {
//	head := CreateLinkedList()
//	reverseLinkedList(head)
//	printLinkedList(head)
//}
//
//func printLinkedList(head *linkedListNode) {
//	itr := head.next
//	for itr != nil {
//		fmt.Printf("%d ", itr.val)
//		itr = itr.next
//	}
//	fmt.Println()
//}
//
//func reverseLinkedList(head *linkedListNode) {
//	previous := head
//	itr := head.next
//	first := itr
//	for itr != nil {
//		next := itr.next
//		current := itr
//		current.next = previous
//		previous = current
//		itr = next
//	}
//	head.next = previous
//	first.next = nil
//}
