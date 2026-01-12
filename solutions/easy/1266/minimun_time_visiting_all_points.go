package mintimetovisitallpoints

func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	seconds := 0
	x, y := points[0][0], points[0][1]
	for i := 1; i < len(points); i++ {
		raw := points[i]

		mx := 0

		if raw[0] > x {
			mx = max(mx, raw[0]-x)
		}

		if raw[0] < x {
			mx = max(mx, x-raw[0])

		}

		if raw[1] > y {
			mx = max(mx, raw[1]-y)
		}

		if raw[1] < y {
			mx = max(mx, y-raw[1])
		}
		x, y = raw[0], raw[1]
		seconds += mx
	}

	return seconds
}
