package main

import (
	"fmt"
	"moneybag/ex00"
)

func resultPrinter(name string, tar int, options []int, res []int) {
	fmt.Printf("%s:|\tTAR: %v\t|\tOPTS: %v\t|\tres: %v\n", name, tar, options, res)
}

func main() {
	tar := 213
	options := []int{3, 100, 10, 5, 1}
	optsForMine := make([]int, len(options))
	copy(optsForMine, options)

	resGiven := ex00.MinCoins(tar, options)
	resMine := ex00.MinCoins2(tar, optsForMine)
	resultPrinter("GIVEN", tar, options, resGiven)
	resultPrinter("MINE", tar, options, resMine)
}
