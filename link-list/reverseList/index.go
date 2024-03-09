package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 开辟一个新的链表来存储数据
// func reverseList(head *ListNode) *ListNode {
// 	var newList *ListNode = nil
// 	for head != nil {
// 		if newList == nil {
// 			newList = &ListNode{}
// 			newList.Val = head.Val
// 		} else {
// 			temp := &ListNode{
// 				Val:  head.Val,
// 				Next: newList,
// 			}
// 			newList = temp
// 		}
// 		head = head.Next
// 	}
// 	return newList
// }

// 2. 在原有的链表上操作
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preNode := head
	currentNode := preNode.Next
	preNode.Next = nil
	var nextNode *ListNode = nil
	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = preNode
		preNode = currentNode
		currentNode = nextNode
	}
	return preNode
}
