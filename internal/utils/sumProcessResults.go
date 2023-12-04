package utils

import (
	"strings"
	"sync"
)

type process func(input string, wg *sync.WaitGroup, resultChan chan<- int)

func SumProcessResults(input string, p process) int {
	lines := strings.Split(input, "\n")

	var wg sync.WaitGroup
	resultChan := make(chan int)

	for _, line := range lines {
		wg.Add(1)
		go p(line, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return SumIntC(resultChan)
}
