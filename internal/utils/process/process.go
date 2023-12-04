package process

import (
	"strings"
	"sync"

	Sum "github.com/TomParfitt/aoc2023/internal/utils/sum"
)

type process func(input string, wg *sync.WaitGroup, resultChan chan<- int)

func SumResults(input string, p process) int {
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

	return Sum.IntC(resultChan)
}
