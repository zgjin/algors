package algors

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [3,9,20,15,7]
// [9,3,15,20,7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}

	rootNode := &TreeNode{
		Val: preorder[0],
	}

	var rootIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
			break
		}
	}
	var inLeftNodes []int
	var inRightNodes []int
	for i := 0; i < rootIndex; i++ {
		inLeftNodes = append(inLeftNodes, inorder[i])
	}
	for i := rootIndex + 1; i < len(inorder); i++ {
		inRightNodes = append(inRightNodes, inorder[i])
	}

	var preLeftNodes []int
	var preRightNodes []int
	for i := 1; i <= len(inLeftNodes); i++ {
		preLeftNodes = append(preLeftNodes, preorder[i])
	}
	for i := 1 + len(inLeftNodes); i < len(preorder); i++ {
		preRightNodes = append(preRightNodes, preorder[i])
	}

	rootNode.Left = buildTree(preLeftNodes, inLeftNodes)
	rootNode.Right = buildTree(preRightNodes, inRightNodes)

	return rootNode
}
