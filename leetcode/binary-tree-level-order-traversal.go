/*
- Problem: https://leetcode.com/problems/binary-tree-level-order-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    
    var queue []*TreeNode = []*TreeNode{}
    var firstNodeOfNextLevel *TreeNode
    var res [][]int = [][]int{}
    
    queue = append(queue, root)
    firstNodeOfNextLevel = root
    
    tmp := []int{}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:len(queue)]
        
        if node != nil {
            // If the current node matches with the firstNodeOfNextLevel it means that we just get into the a new level of the tree.
            // So we can stop pushing items to tmp array now. And push the tmp array to the result array as it now contains all the values of the nodes of the previous level
            if node == firstNodeOfNextLevel {
                if len(tmp) > 0 {
                    res = append(res, tmp)
                    tmp = []int{}
                }
                firstNodeOfNextLevel = nil
            }
            
            tmp = append(tmp, node.Val)
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
            
            if firstNodeOfNextLevel == nil {
                firstNodeOfNextLevel = node.Left
                if firstNodeOfNextLevel == nil {
                    firstNodeOfNextLevel = node.Right
                }
            }
        }
        
    }
    
    if len(tmp) > 0 {
        res = append(res, tmp)
    }
    
    return res
    
}
