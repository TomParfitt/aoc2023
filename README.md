# Advent Of Code 
https://adventofcode.com/2023

## Recepie Book
Some common snippets ...

### Unbounded GoRoutines
```go
func unboundedExample() int {

	var wg sync.WaitGroup
	resultChan := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go process(i, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for r := range resultChan {
		sum += r
	}

	return sum
}

func process(i int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	resultChan <- i
}
```
