package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// ?
	res := []float64{}
	n := len(costs)
	for i := 0; i < n; i++ {
		if costs[i].day == day {
			res = append(res, costs[i].value)
		}
	}
	return res
}
