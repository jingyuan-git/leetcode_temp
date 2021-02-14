//Given the root of a binary tree, flatten the tree into a "linked list": 
//
// 
// The "linked list" should use the same TreeNode class where the right child po
//inter points to the next node in the list and the left child pointer is always n
//ull. 
// The "linked list" should be in the same order as a pre-order traversal of the
// binary tree. 
// 
//
// 
// Example 1: 
//
// 
//Input: root = [1,2,5,3,4,null,6]
//Output: [1,null,2,null,3,null,4,null,5,null,6]
// 
//
// Example 2: 
//
// 
//Input: root = []
//Output: []
// 
//
// Example 3: 
//
// 
//Input: root = [0]
//Output: [0]
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the tree is in the range [0, 2000]. 
// -100 <= Node.val <= 100 
// 
//
// 
//Follow up: Can you flatten the tree in-place (with O(1) extra space)? Related 
//Topics 树 深度优先搜索 
// 👍 720 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		pre, curr := list[i-1], list[i]
		pre.Left, pre.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
//leetcode submit region end(Prohibit modification and deletion)
