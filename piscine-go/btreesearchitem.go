package piscine

type TreeNode4 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData1(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	current := root
	var parent *TreeNode

	for current != nil {
		parent = current
		if data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	newNode4 := &TreeNode{Data: data, Parent: parent}
	if data < parent.Data {
		parent.Left = newNode4
	} else {
		parent.Right = newNode4
	}

	return root
}
