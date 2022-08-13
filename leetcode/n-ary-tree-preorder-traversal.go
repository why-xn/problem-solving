/*
- Problem: https://leetcode.com/problems/n-ary-tree-preorder-traversal/
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    
    var res []int = []int{}
    dfs(root, &res)
    return res
}

func dfs(node *Node, res *[]int) {
    if node == nil {
        return
    }
    *res = append(*res, node.Val)
    for _, c := range node.Children {
        dfs(c, res)
    }
}
