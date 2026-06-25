package createbinarytreefromdesc

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeStruct struct {
	hasParent bool
	TreeNode  *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {

	m := make(map[int]TreeNodeStruct)

	for i := 0; i < len(descriptions); i++ {
		node := descriptions[i]
		parentTreeStruct, ok := m[node[0]]
		parentTree := parentTreeStruct.TreeNode
		if ok {
			setChild(parentTree, m, node)
			continue
		} else {
			parentTree = &TreeNode{Val: node[0]}
			m[node[0]] = TreeNodeStruct{false, parentTree}
		}

		setChild(parentTree, m, node)
	}

	for _, nodeStruct := range m {
		if nodeStruct.hasParent == false {
			return nodeStruct.TreeNode
		}
	}

	return &TreeNode{}
}

func setChild(parentTree *TreeNode, m map[int]TreeNodeStruct, node []int) {
	child, chOk := m[node[1]]
	if node[2] == 1 {
		if chOk {
			parentTree.Left = child.TreeNode
		} else {
			parentTree.Left = &TreeNode{Val: node[1]}
		}
		m[node[1]] = TreeNodeStruct{true, parentTree.Left}
	} else {
		if chOk {
			parentTree.Right = child.TreeNode
		} else {
			parentTree.Right = &TreeNode{Val: node[1]}
		}
		m[node[1]] = TreeNodeStruct{true, parentTree.Right}
	}
}
