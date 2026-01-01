package __connections_asm

func countConnections(groupSize int) int {
	// ?
	total := 0
	for i := 1; i <= groupSize; i++ {
		for j := 1; j <= groupSize-1; j++ {
			total++
		}
	}
	return total / 2
}
