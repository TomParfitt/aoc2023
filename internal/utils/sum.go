package utils

func SumIntS(slice []int) int {
	sum := 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func SumIntC(resultChan chan int) int {
	sum := 0
	for s := range resultChan {
		sum += s
	}
	return sum
}
