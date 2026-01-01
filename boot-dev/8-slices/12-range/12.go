package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?
	for i, value := range msg {
		for _, badWord := range badWords {
			if value == badWord {
				return i
			}
		}
	}
	return -1
}
