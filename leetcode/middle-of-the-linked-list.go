/*
- Problem: https://leetcode.com/problems/middle-of-the-linked-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    
    secondPointer := head
    
    for secondPointer != nil && secondPointer.Next != nil {
        // Increasing Second Pointer upto 2 times in each iteration. When the Second Pointer will reach the last node, The head pointer will be in the middle.
        secondPointer = secondPointer.Next.Next
        head = head.Next
    }
    
    return head
    
}
