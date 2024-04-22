package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readNum() (int, bool) {
	var inp string
	if _, err := fmt.Scanf("%s\n", &inp); err != nil {
		return 0, true
	}
	res, err := strconv.Atoi(inp)
	if err != nil || res < MinLimit || res > MaxLimit {
		fmt.Printf(
			"error [input]: iput should be integer and %d > n < %d\n",
			MinLimit,
			MaxLimit,
		)
		os.Exit(1)
	}

	return res, false
}

func initializeNumList() []int {
	var nums []int
	for {
		num, finished := readNum()
		if finished {
			break
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	return nums
}

func printStats(stats map[string]float64, options []string) {
	for _, name := range options {
		fmt.Printf("%s: %.2f\n", name, stats[name])
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func initializeOptions(variants []string) []string {
	options := flag.String(
		"options",
		strings.Join(variants, "+"),
		"Pick the options you need. Available options are: Mean, Mode, SD, Median. e.g: -options Median\nJoin the options with a '+' to pick multiple. e.g: -options Mean+SD",
	)
	var res []string
	flag.Parse()
	if len(*options) > 0 {
		for _, option := range strings.Split(*options, "+") {
			if !contains(variants, option) {
				fmt.Printf("Options should be one of: %v\n", variants)
				os.Exit(1)
			}
			res = append(res, option)
		}
	}
	return res
}

func main() {
	var nums []int
	var statistics map[string]float64
	statisticsOptions := initializeOptions(DefaultStatOptions)

	nums = initializeNumList()
	statistics = GetStatistics(nums)
	printStats(statistics, statisticsOptions)
}
