package toysontree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SumToysCase struct {
	name   string
	args   *TreeNode
	expect int
}
type AreToysBalancedCase struct {
	name   string
	args   *TreeNode
	expect bool
}

func TestAreToysBalanced(t *testing.T) {
	balanced := &TreeNode{
		HasToy: false,
		Left:   NewTreeNode(true).AddRight(false).AddLeft(false),
		Right: &TreeNode{
			Left:  NewTreeNode(false).AddLeft(false).AddRight(true),
			Right: NewTreeNode(false).AddRight(false),
		}}
	notBalanced := &TreeNode{
		HasToy: false,
		Left:   NewTreeNode(true).AddRight(false).AddLeft(false),
		Right: &TreeNode{
			Left:  NewTreeNode(false).AddLeft(false).AddRight(true),
			Right: NewTreeNode(false).AddRight(false).AddLeft(true),
		}}

	cases := []AreToysBalancedCase{
		{
			name:   "Should balanced be true",
			args:   balanced,
			expect: true,
		},
		{
			name:   "Should not balanced be false",
			args:   notBalanced,
			expect: false,
		},
		{
			name:   "Root should be ingored, and empty trees (root only) should be always balanced",
			args:   NewTreeNode(false),
			expect: true,
		},
		{
			name:   "Root should be ingored, and empty trees (root only) should be always balanced",
			args:   NewTreeNode(true),
			expect: true,
		},
		{
			name:   "Single side should be false",
			args:   NewTreeNode(false).AddLeft(true),
			expect: false,
		},
		{
			name:   "Both sides false should be true",
			args:   NewTreeNode(false).AddLeft(false).AddRight(false),
			expect: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := AreToysBalanced(c.args)
			assert.Equal(t, c.expect, got)
		})
	}

}

func TestSumToys(t *testing.T) {
	equalsThree := &TreeNode{
		HasToy: false,
		Left:   NewTreeNode(true).AddRight(false).AddLeft(false),
		Right: &TreeNode{
			Left:  NewTreeNode(false).AddLeft(false).AddRight(true),
			Right: NewTreeNode(false).AddRight(true),
		}}

	equalsFour := &TreeNode{
		HasToy: false,
		Left:   NewTreeNode(true).AddRight(false).AddLeft(true),
		Right: &TreeNode{
			Left:  NewTreeNode(false).AddLeft(false).AddRight(true),
			Right: NewTreeNode(false).AddRight(false).AddLeft(true),
		}}
	empty := &TreeNode{HasToy: false}
	singleDeepSide := &TreeNode{
		HasToy: false,
		Left: &TreeNode{
			HasToy: false,
			Left: &TreeNode{
				HasToy: false,
				Left:   NewTreeNode(false).AddLeft(true),
			},
		},
	}

	cases := []SumToysCase{
		{
			name:   "Empty tree",
			args:   empty,
			expect: 0,
		},
		{
			name:   "Single left node",
			args:   NewTreeNode(false).AddLeft(true),
			expect: 1,
		},
		{
			name:   "Single right node",
			args:   NewTreeNode(false).AddRight(true),
			expect: 1,
		},
		{
			name:   "Single deep left side",
			args:   singleDeepSide,
			expect: 1,
		},
		{
			name:   "Should equal three",
			args:   equalsThree,
			expect: 3,
		},
		{
			name:   "Should equal four",
			args:   equalsFour,
			expect: 4,
		},
		{
			name:   "No manual assign boolean",
			args:   &TreeNode{},
			expect: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := SumToys(c.args)
			assert.Equal(t, c.expect, got)
		})
	}
}
