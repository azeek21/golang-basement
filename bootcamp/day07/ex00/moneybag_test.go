package ex00_test

import (
	"fmt"
	"moneybag/ex00"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Desc    string
	Target  int
	Options []int
	Exp     []int
}

func getCases() []TestCase {

	return []TestCase{
		{
			Desc:    "given by task",
			Target:  13,
			Options: []int{1, 5, 10},
			Exp:     []int{10, 1, 1, 1},
		},
		{
			Desc:    "unordered options",
			Target:  13,
			Options: []int{1, 10, 5},
			Exp:     []int{10, 1, 1, 1},
		},
		{
			Desc:    "duplicate options",
			Target:  13,
			Options: []int{1, 1, 5, 5, 5, 10},
			Exp:     []int{10, 1, 1, 1},
		},
		{
			Desc:    "empty options",
			Target:  13,
			Options: []int{},
			Exp:     []int{},
		},
		{
			Desc:    "duplicate unordered options",
			Target:  13,
			Options: []int{10, 5, 1, 5, 5, 5, 10, 1, 5, 1, 10, 1},
			Exp:     []int{10, 1, 1, 1},
		},
		{
			Desc:    "empty",
			Target:  13,
			Options: []int{},
			Exp:     []int{},
		},
		{
			Desc:    "not enough options",
			Target:  13,
			Options: []int{4, 10},
			Exp:     []int{10},
		},
		{
			Desc:    "single target",
			Target:  1,
			Options: []int{1, 4, 10},
			Exp:     []int{1},
		},
		{
			Desc:    "single target options mixed",
			Target:  1,
			Options: []int{4, 1, 10},
			Exp:     []int{1},
		},
		{
			Desc:    "target=2 target options mixed",
			Target:  2,
			Options: []int{4, 1, 10},
			Exp:     []int{1, 1},
		},
		{
			Desc:    "target=2 target options mixed",
			Target:  2,
			Options: []int{4, 2, 1, 10},
			Exp:     []int{2},
		},
	}

}

func announce(stage, name string) string {
	return fmt.Sprintf("**********\t%s: %s\t**********\n", stage, name)
}

func asdf(t *testing.T) {
	// Arrange
	testCases := getCases()

	t.Log(announce("Failed: ", "minCoins (given)"))
	// Act
	for n, tCase := range testCases {
		t.Run(fmt.Sprintf("%d-%s", n+1, tCase.Desc), func(t *testing.T) {
			actual := ex00.MinCoins(tCase.Target, tCase.Options)
			// assert
			assert.ElementsMatch(t, actual, tCase.Exp)
		})
	}

}

func TestMinCoins2(t *testing.T) {
	// Arrange
	testCases := getCases()

	t.Log(announce("Failed: ", "minCoins2 (mine)"))
	// Act
	for n, tCase := range testCases {
		t.Run(fmt.Sprintf("%d-%s", n+1, tCase.Desc), func(t *testing.T) {
			actual := ex00.MinCoins2(tCase.Target, tCase.Options)
			// assert
			assert.ElementsMatch(t, actual, tCase.Exp)
		})
	}
}

func ExampleMinCoins() {
	ex00.MinCoins2(13, []int{5, 10, 1})
	//Output: 1,1,1,1,1,1,1,1,1,1, 1, 1, 1 (13 * 1)
}

func ExampleMinCoins2() {
	ex00.MinCoins2(13, []int{5, 10, 1})
	//Output: 10, 1, 1,
}
