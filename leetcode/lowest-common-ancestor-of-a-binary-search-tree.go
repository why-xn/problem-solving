/*
- Problem: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type HashSet struct {
	set map[int]*byte
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var ancestorSetOfP HashSet = HashSet{set: map[int]*byte{}}    // For storing ancestors of P in bottom up order
	var ancestorSetOfQ HashSet = HashSet{set: map[int]*byte{}}    // For storing ancestors of Q in bottom up order
	var lca TreeNode = TreeNode{
		Val:   -1000000001,
		Right: nil,
		Left:  nil,
	}

	dfs(root, p, q, &ancestorSetOfP, &ancestorSetOfQ, &lca)
	return &lca

}

func dfs(node, p, q *TreeNode, ancestorSetOfP *HashSet, ancestorSetOfQ *HashSet, lca *TreeNode) {
	if node != nil {
		if node.Val == p.Val {
			ancestorSetOfP.set[node.Val] = nil
		} else if node.Val == q.Val {
			ancestorSetOfQ.set[node.Val] = nil
		}
		if lca.Val == -1000000001 {
			dfs(node.Left, p, q, ancestorSetOfP, ancestorSetOfQ, lca)
			checkIfAncestorOfBoth(node, ancestorSetOfP, ancestorSetOfQ, lca)

			dfs(node.Right, p, q, ancestorSetOfP, ancestorSetOfQ, lca)
			checkIfAncestorOfBoth(node, ancestorSetOfP, ancestorSetOfQ, lca)
		}
	}
}

func checkIfAncestorOfBoth(node *TreeNode, ancestorSetOfP *HashSet, ancestorSetOfQ *HashSet, lca *TreeNode) {
	isAncestorOfP := isAncestor(node, ancestorSetOfP)
	isAncestorOfQ := isAncestor(node, ancestorSetOfQ)

	// If the node is ancestor of both P and Q, then this is the Lowest Common Ancestor.
	if isAncestorOfP && isAncestorOfQ {
		if lca.Val == -1000000001 {
			*lca = *node
		}
	}
}

func isAncestor(node *TreeNode, ancestorSet *HashSet) bool {
	isAncestor := false

	// checking if left child is an ancestor. If so, this node is also an ancestor
	if node.Left != nil {
		if _, exists := ancestorSet.set[node.Left.Val]; exists {
			isAncestor = true
		}
	}

	// checking if right child is an ancestor. If so, this node is also an ancestor
	if !isAncestor && node.Right != nil {
		if _, exists := ancestorSet.set[node.Right.Val]; exists {
			isAncestor = true
		}
	}

	// checking if this node is already in ancestor.
	if !isAncestor {
		if _, exists := ancestorSet.set[node.Val]; exists {
			isAncestor = true
		}
	}

	if isAncestor {
		if _, exists := ancestorSet.set[node.Val]; !exists {
			ancestorSet.set[node.Val] = nil
		}
	}

	return isAncestor
}
