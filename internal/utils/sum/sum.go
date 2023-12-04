package sum

func IntS(slice []int) int {
	sum := 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func IntC(resultChan chan int) int {
	sum := 0
	for s := range resultChan {
		sum += s
	}
	return sum
}
