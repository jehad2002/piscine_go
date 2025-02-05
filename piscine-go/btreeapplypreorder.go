package piscine

type TreeNode3 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeApplyPreorder applies a given function to each element in the tree using preorder traversal
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Apply the function to the current node
	f(root.Data)

	// Traverse the left subtree
	BTreeApplyPreorder(root.Left, f)

	// Traverse the right subtree
	BTreeApplyPreorder(root.Right, f)
}
