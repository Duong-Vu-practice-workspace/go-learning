package main

import "slices"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?
	for i, value := range msg {
		if slices.Contains(badWords, value) {
				return i
			}
	}
	return -1
}
