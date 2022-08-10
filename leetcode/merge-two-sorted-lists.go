/*
- Problem: https://leetcode.com/problems/merge-two-sorted-lists/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    } else if list1 == nil && list2 == nil {
        return nil
    }
    
    // Determining the top node of the result.
    result := list1
    var tmp *ListNode
    
    if list1.Val > list2.Val {
        result = list2
        list2 = list2.Next
    } else {
        list1 = list1.Next
    }
    
    tmp = result    
    
    for {  
        if list1 == nil {
            tmp.Next = list2
            break
        } else if list2 == nil {
            tmp.Next = list1
            break
        }
        
        // Determining which top node between the two linked list is the smaller one and push that to the result
        if list1.Val <= list2.Val {
            tmp.Next = list1
            list1 = list1.Next
        } else {
            tmp.Next = list2
            list2 = list2.Next
        }
        tmp = tmp.Next
    }
    return result
}
