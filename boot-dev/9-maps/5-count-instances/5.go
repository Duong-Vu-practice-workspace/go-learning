package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	// ?
	for i := range messagedUsers {
		_, ok := validUsers[messagedUsers[i]]
		if ok {
			validUsers[messagedUsers[i]]++
		}
	}
}
