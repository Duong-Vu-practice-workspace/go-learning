package __omitting_conditions

func maxMessages(thresh int) int {
	// ?
	totalFee := 0
	for i := 0; ; i++ {
		totalFee += 100 + i
		if totalFee > thresh {
			return i
		}
	}
}
