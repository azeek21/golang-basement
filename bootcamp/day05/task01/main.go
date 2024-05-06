package main

import (
	toysontree "day05/task00/toys-on-tree"
	"day05/utils"
	"fmt"
)

// breath first traverse (BFS)
func BFSFlat(root *toysontree.TreeNode) [][]bool {

	if root == nil {
		return nil
	}

	res := [][]bool{}
	que := utils.NewQueue()
	que.Enque(root)

	for !que.Empty() {
		tmp := []bool{}
		for size := que.Len(); size > 0; size-- {
			_item, _ := que.Deque()
			item := _item.(*toysontree.TreeNode)
			tmp = append(tmp, item.HasToy)
			if item.Left != nil {
				que.Enque(item.Left)
			}
			if item.Right != nil {
				que.Enque(item.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}

/*
		   1
		  /  \
		 1     0
		/ \   / \
	       1   0 1   1
*/
func main() {
	tree := toysontree.TreeNode{
		HasToy: true,
		Left:   toysontree.NewTreeNode(true).AddLeft(true).AddRight(false),
		Right:  toysontree.NewTreeNode(false).AddLeft(true).AddRight(true),
	}

	res := BFSFlat(&tree)
	fmt.Printf("RES: %v\n", res)
}
