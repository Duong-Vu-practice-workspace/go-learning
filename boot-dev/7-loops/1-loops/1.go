package __loops

func bulkSend(numMessages int) float64 {
	// ?
	var totalCost float64
	for i := 0; i < numMessages; i++ {
		x := float64(i)
		x *= 1.0
		stepResult := 1.0 + (x / 100.0)
		totalCost += stepResult
	}
	return totalCost
}
