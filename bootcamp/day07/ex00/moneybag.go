// This is where all essential functions are stored
// ex00 mainly focuses on fixing problems of provided MinCoins function and writing unit tests
package ex00

// This function was provided as an example to test
// and find it's problems. It's problems are fixed at MinCoins2
// BUG: input: 10, 5, 1 | output: 1, 1, 1, 1, 1, 1, 1, 1, 1, 1 | expected: 10, 1, 1
func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// This is a version of MinCoins but with fixed bugs.
// It fixes problme where MinCoins doesn't choose the
// most optimal value in terms of unsorted array.
func MinCoins2(target int, options []int) []int {
	QuickSort(options)

	res := make([]int, 0)
	i := len(options) - 1

	for i >= 0 {
		for target >= options[i] {
			target -= options[i]
			res = append(res, options[i])
		}
		i -= 1
	}

	return res
}

// MinCoins2Optimized is a more performance oriented version of MinCoins2Optimized
// It reajes almost 2-3 times faster competions just by avoiding sorting the array if it's already sorted
func MinCoins2Optimized(target int, options []int) []int {
	if !IsSorted(options) {
		QuickSort(options)
	}

	res := make([]int, 0)
	i := len(options) - 1

	for i >= 0 {
		for target >= options[i] {
			target -= options[i]
			res = append(res, options[i])
		}
		i -= 1
	}

	return res
}
