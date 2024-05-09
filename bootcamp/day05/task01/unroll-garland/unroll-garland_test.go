package unrollgarland_test

import (
	toysontree "day05/task00/toys-on-tree"
	unrollgarland "day05/task01/unroll-garland"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseUnlrollGrand struct {
	name   string
	arg    *toysontree.TreeNode
	expect []bool
}

func TestUnrollGarland(t *testing.T) {
	/*
			   1
			  /  \
			 1     0
			/ \   / \
		       1   0 1   1
	*/
	givenTreeUnrolled := []bool{true, true, false, true, true, false, true}
	givenTree := toysontree.TreeNode{
		HasToy: true,
		Left:   toysontree.NewTreeNode(true).AddLeft(true).AddRight(false),
		Right:  toysontree.NewTreeNode(false).AddLeft(true).AddRight(true),
	}

	/*
	    0
	   / \
	  0   1
	*/
	basicTreeUnrolled := []bool{false, false, true}
	basicTtree := toysontree.TreeNode{
		HasToy: false,
		Left:   toysontree.NewTreeNode(false),
		Right:  toysontree.NewTreeNode(true),
	}

	/*
			         0
			    /          \
			   1            0
			 /   \        /   \
		        0     1      1     0
		       / \   / \    / \   / \
		      0   1 1   0  1   0 0   1
	*/

	deepTreeUnrolled := []bool{false,
		true, false,
		false, true, true, false,
		false, true, true, false, true, false, false, true}
	deepTree := toysontree.TreeNode{
		HasToy: false,
		Left: &toysontree.TreeNode{
			HasToy: true,
			Left:   toysontree.NewTreeNode(false).AddLeft(false).AddRight(true),
			Right:  toysontree.NewTreeNode(true).AddLeft(true).AddRight(false),
		},
		Right: &toysontree.TreeNode{
			HasToy: false,
			Left:   toysontree.NewTreeNode(true).AddLeft(true).AddRight(false),
			Right:  toysontree.NewTreeNode(false).AddLeft(false).AddRight(true),
		},
	}

	testCases := []TestCaseUnlrollGrand{
		{
			name:   "Given tree in task",
			arg:    &givenTree,
			expect: givenTreeUnrolled,
		},
		{
			name:   "Basec tree",
			arg:    &basicTtree,
			expect: basicTreeUnrolled,
		},
		{
			name:   "Nested tree",
			arg:    &deepTree,
			expect: deepTreeUnrolled,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := unrollgarland.UnrollGarland(testCase.arg)
			assert.Equal(t, testCase.expect, res)
		})

	}
}
