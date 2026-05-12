package minimuminitialenery

// Title: 1665. Minimum Initial Energy to Finish Tasks
// Link: https://leetcode.com/problems/minimum-initial-energy-to-finish-tasks/

import "sort"

func minimumEffort(tasks [][]int) int {
	res := 0

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] < tasks[j][1]-tasks[j][0]
	})

	for _, task := range tasks {
		res = max(res+task[0], task[1])
	}

	return res
}
