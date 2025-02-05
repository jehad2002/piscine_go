package piscine

type TreeNode2 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyPostorder(root.Left, f)

	BTreeApplyPostorder(root.Right, f)

	f(root.Data)
}
