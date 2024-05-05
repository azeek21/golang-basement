package main

import (
	toysontree "day05/task00/toys-on-tree"
	"log"
)

func main() {
	balancedToys := &toysontree.TreeNode{
		HasToy: false,
		Left:   toysontree.NewTreeNode(true).AddRight(true).AddLeft(false),
		Right: &toysontree.TreeNode{
			Left:  toysontree.NewTreeNode(false).AddLeft(false).AddRight(true),
			Right: toysontree.NewTreeNode(false).AddRight(true),
		}}

	notBalancedToys := &toysontree.TreeNode{
		HasToy: false,
		Left:   toysontree.NewTreeNode(true).AddRight(false).AddLeft(false),
		Right: &toysontree.TreeNode{
			Left:  toysontree.NewTreeNode(false).AddLeft(false).AddRight(true),
			Right: toysontree.NewTreeNode(false).AddRight(true),
		}}

	log.Printf("balancedToys is balanced: %v\n", toysontree.AreToysBalanced(balancedToys))

	log.Printf("notBalancedToys is balanced: %v\n", toysontree.AreToysBalanced(notBalancedToys))

}
