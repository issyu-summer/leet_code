package _024_12_16

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lastInterval := merged[len(merged)-1]
		//try to merge
		if lastInterval[1] >= intervals[i][0] {
			merged[len(merged)-1][1] = max(lastInterval[1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
