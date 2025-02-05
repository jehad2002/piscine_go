package piscine

type TreeNode1 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyInorder(root.Left, f)

	f(root.Data)

	BTreeApplyInorder(root.Right, f)
}
