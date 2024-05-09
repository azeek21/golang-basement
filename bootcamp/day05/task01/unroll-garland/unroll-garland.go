package unrollgarland

import (
	toysontree "day05/task00/toys-on-tree"
	"day05/utils"
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

func UnrollGarland(root *toysontree.TreeNode) []bool {
	treeMatrix := BFSFlat(root)
	var res []bool
	for layerIndex, row := range treeMatrix {
		if layerIndex%2 != 0 {
			res = append(res, row...)
		} else {
			for i := len(row) - 1; i >= 0; i-- {
				res = append(res, row[i])
			}
		}
	}
	return res
}
