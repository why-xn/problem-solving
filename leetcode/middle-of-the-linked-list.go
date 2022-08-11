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
    var i int = 0
    
    for secondPointer.Next != nil {
        i = 0
      
        // Increasing Second Pointer upto 2 times in each iteration. When the Second Pointer will reach the last node, The head pointer will be in the middle.
        for i < 2 {
            if secondPointer.Next != nil {
                secondPointer = secondPointer.Next
            }
            i++
        }
        head = head.Next
    }
    
    return head
    
}
