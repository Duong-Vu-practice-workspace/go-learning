package __arrays

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	// ?
	var res1 [3]string
	var res2 [3]int
	res1 = [3]string{primary, secondary, tertiary}
	res2 = [3]int{len(primary), len(secondary) + len(primary), len(primary) + len(secondary) + len(tertiary)}
	return res1, res2
}
