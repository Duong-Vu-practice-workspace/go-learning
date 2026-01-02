package main

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	res := make(map[rune]map[string]int)
	for i := range names {
		runes := []rune(names[i])
		firstChar := runes[0]
		name := names[i]
		if _, ok := res[firstChar]; !ok {
			res[firstChar] = make(map[string]int)
		}
		res[firstChar][name]++
	}
	return res
}
