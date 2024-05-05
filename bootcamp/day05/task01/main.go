package main

import (
	toysontree "day05/task00/toys-on-tree"
	"log"
)

// breath first traverse (BFS)
func BFSFlat(root *toysontree.TreeNode) [][]bool {
	if root == nil {
		return nil
	}

	ans := [][]bool{}
	queue := []*toysontree.TreeNode{root}

	for len(queue) > 0 {
		tmp := []bool{}
		log.Printf("Q LEN: %v\n", len(queue))
		for i, node := range queue {
			log.Printf("N I: %v\n", i)
			queue = queue[1:]
			tmp = append(tmp, node.HasToy)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, tmp)
	}

	return ans
}

/*
		   1
		  /  \
		 1     0
		/ \   / \
	       1   0 1   1
*/

func UnrollGarland(root *toysontree.TreeNode) []bool {
	tree := toysontree.TreeNode{
		HasToy: true,
		Left:   toysontree.NewTreeNode(true).AddLeft(true).AddRight(false),
		Right:  toysontree.NewTreeNode(false).AddLeft(true).AddRight(true),
	}
	flat := BFSFlat(&tree)
	res := []bool{}
	for i, layer := range flat {
		if i%2 != 0 {
			res = append(res, layer...)
		} else {
			for j := len(layer) - 1; j >= 0; j-- {
				res = append(res, layer[j])
			}
		}
	}
	return res
}

func main() {
	res := UnrollGarland(nil)
	log.Printf("%v\n", res)
}
