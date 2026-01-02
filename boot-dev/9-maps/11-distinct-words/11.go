package main

import "strings"

func countDistinctWords(messages []string) int {
	// ?
	res := make(map[string]struct{})
	for i := range messages {
		words := strings.Split(strings.ToLower(messages[i]), " ")
		for j := range words {
			word := words[j]
			if word != "" {
				if _, ok := res[word]; !ok {
					res[word] = struct{}{}
				}
			}

		}
	}
	return len(res)
}
