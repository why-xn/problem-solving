/*
- Problem: https://leetcode.com/problems/linked-list-cycle-ii/
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 
func detectCycle(head *ListNode) *ListNode {
    
    var hash map[*ListNode]*byte = map[*ListNode]*byte{}
    
    for head != nil {
        // Checking if the node exists in the hash map. If exists, then it was already traversed and it is the starting point of the cycle.
        if _, ok := hash[head]; ok {
            return head
        } else {
            hash[head] = nil
            head = head.Next
        }
    }
    
    return nil
    
}
