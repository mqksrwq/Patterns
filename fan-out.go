// Fan-out - паттерн, позволяющий распараллелить задачу, то есть
// разделить ее на несколько горутин

package Patterns

import "fmt"

func generatorFO(nums []*int, doneCh chan struct{}) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)
		for _, num := range nums {
			select {
			case <-doneCh:
				return
			case resultCh <- *num:
			}
		}
	}()
	return resultCh
}

func addFO(inputCh chan int, doneCh chan struct{}) chan int {
	resultCh := make(chan int)
	go func() {
		defer close(resultCh)
		for num := range inputCh {
			result := num + 1
			select {
			case <-doneCh:
				return
			case resultCh <- result:
			}
		}
	}()
	return resultCh
}

func FanOut(inputCh chan int, doneCh chan struct{}, workers int) []chan int {
	resultChs := make([]chan int, workers)
	for i := 0; i < workers; i++ {
		resultChs[i] = addFO(inputCh, doneCh)
	}
	return resultChs
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doneCh := make(chan struct{})
	defer close(doneCh)
	ch := generator(nums)
	addCh := add(ch, doneCh)
	for num := range addCh {
		fmt.Println(num)
	}
}
