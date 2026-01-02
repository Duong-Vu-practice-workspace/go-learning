package main

import "strings"

func removeProfanity(message *string) {
	// ?
	content := *message
	content = strings.ReplaceAll(content, "fubb", "****")
	content = strings.ReplaceAll(content, "shiz", "****")
	content = strings.ReplaceAll(content, "witch", "*****")

	*message = content
}
