package main

func addEmailsToQueue(emails []string) chan string {
	// ?
	n := len(emails)
	channel := make(chan string, n)
	for i := range emails {
		channel <- emails[i]
	}
	return channel
}
