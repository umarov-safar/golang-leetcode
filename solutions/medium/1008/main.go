package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return &TreeNode{}
	}

	var tree = &TreeNode{Val: preorder[0]}

	for i := 1; i < len(preorder); i++ {
		num := preorder[i]
		var currentNode = tree
		for {
			if currentNode.Val > num {
				if currentNode.Left == nil {
					currentNode.Left = &TreeNode{Val: num}
					break
				}

				currentNode = currentNode.Left
			}

			if currentNode.Val < num {
				if currentNode.Right == nil {
					currentNode.Right = &TreeNode{Val: num}
					break
				}

				currentNode = currentNode.Right
			}
		}
	}

	return tree
}
