/*
- Problem: https://leetcode.com/problems/reverse-linked-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
	var newHead ListNode
	reverse(head, &newHead)
	return &newHead
}

func reverse(node *ListNode, newHead *ListNode) *ListNode {
	if node.Next == nil {
		*newHead = *node
		return newHead
	}
	next := reverse(node.Next, newHead)
  
  // Reversing the order
	next.Next = node
	node.Next = nil
	return node
}
