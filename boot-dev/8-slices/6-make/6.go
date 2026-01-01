package __make

func getMessageCosts(messages []string) []float64 {
	// ?
	n := len(messages)
	res := make([]float64, n)

	for i := 0; i < n; i++ {
		res[i] = float64(len(messages[i])) * 0.01
	}
	return res
}
