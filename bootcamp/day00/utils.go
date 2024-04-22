package main

import "math"

func getMedian(nums []int) float64 {
	length := len(nums)
	if length%2 == 0 {
		return float64(nums[length/2-1]+nums[length/2]) / 2
	}
	return float64(nums[length/2])
}

func getMode(nums []int) int {
	heatMap := map[int]int{}
	mode := nums[0]

	for _, n := range nums {
		heatMap[n]++
	}

	for num, occurrence := range heatMap {
		if occurrence > heatMap[mode] {
			mode = num
		}
	}
	return mode
}

func getMean(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}

func getStandardDeviation(nums []int, mean float64) float64 {
	sd := 0.0
	for _, n := range nums {
		sd += math.Pow(float64(n)-mean, 2)
	}
	return math.Sqrt(sd / float64(len(nums)))
}

func GetStatistics(nums []int) map[string]float64 {
	stats := map[string]float64{}
	stats["Median"] = getMedian(nums)
	stats["Mode"] = float64(getMode(nums))
	stats["Mean"] = getMean(nums)
	stats["SD"] = getStandardDeviation(nums, stats["mean"])
	return stats
}
