package main

import (
	toysontree "day05/task00/toys-on-tree"
	unrollgarland "day05/task01/unroll-garland"
	"fmt"
)

func main() {
	/*
			   1
			  /  \
			 1     0
			/ \   / \
		       1   0 1   1
	*/
	tree := toysontree.TreeNode{
		HasToy: true,
		Left:   toysontree.NewTreeNode(true).AddLeft(true).AddRight(false),
		Right:  toysontree.NewTreeNode(false).AddLeft(true).AddRight(true),
	}

	res := unrollgarland.UnrollGarland(&tree)
	fmt.Printf("UNROLLED: %v\n", res)
}
