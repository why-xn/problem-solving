/*
- Problem: https://leetcode.com/problems/validate-binary-search-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    
    return isNodeValid(root, 2147483648, -2147483649)
    
}

func isNodeValid(node *TreeNode, upperLimit int, lowerLimit int) bool {
    if node != nil {
        if node.Val >= upperLimit {
            return false
        } else if node.Val <= lowerLimit {
            return false
        }

        return isNodeValid(node.Left, node.Val, lowerLimit) && isNodeValid(node.Right, upperLimit, node.Val)
    }
    return true
}
