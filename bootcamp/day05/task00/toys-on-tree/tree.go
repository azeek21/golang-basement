package toysontree

import (
	"day05/utils"
	"log"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(hasToy bool) *TreeNode {
	return &TreeNode{
		HasToy: hasToy,
	}
}

func (node *TreeNode) AddLeft(hasToy bool) *TreeNode {
	node.Left = NewTreeNode(hasToy)
	return node
}

func (node *TreeNode) AddRight(hasToy bool) *TreeNode {
	node.Right = NewTreeNode(hasToy)
	return node
}

// Depth first travers (DFS)
func SumToys(treeRoot *TreeNode) int {
	// base cases
	if treeRoot == nil {
		return 0
	}
	left := SumToys(treeRoot.Left)
	right := SumToys(treeRoot.Right)
	return utils.Bool2int(treeRoot.HasToy) + left + right
}

// Tree balance check
func AreToysBalanced(root *TreeNode) bool {
	if root == nil {
		log.Println("error: Tree root is null")
		return true
	}
	left := SumToys(root.Left)
	right := SumToys(root.Right)
	return left == right
}
