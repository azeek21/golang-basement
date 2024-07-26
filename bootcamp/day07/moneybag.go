package main

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

func unque(src []int) []int {
	set := map[int]bool{}
	for _, i := range src {
		if !set[i] {
			set[i] = true
		}
	}

	res := make([]int, len(set))
	i := 0
	for key := range set {
		res[i] = key
		i++
	}
	return res
}

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
