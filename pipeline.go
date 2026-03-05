// Pipeline - паттерн многопоточного проектирования позволяющий разбить
// задачу на несколько подзадач, которые выполняются независимо друг от
// друга

package Patterns

import "fmt"

func add(inputCh <-chan int, doneCh <-chan struct{}) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)
		for value := range inputCh {
			result := value + 1

			select {
			case <-doneCh:
				return
			case resultCh <- result:
			}
		}
	}()
	return resultCh
}

func mult(inputCh <-chan int, doneCh <-chan struct{}) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)
		for value := range inputCh {
			result := value * 2

			select {
			case <-doneCh:
				return
			case resultCh <- result:
			}
		}
	}()
	return resultCh
}

func generatorPipeline(doneCh <-chan struct{}, numbers []int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)
		for _, value := range numbers {

			select {
			case <-doneCh:
				return
			case resultCh <- value:
			}
		}
	}()
	return resultCh
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doneCh := make(chan struct{})
	inputCh := generatorPipeline(doneCh, numbers)
	addCh := add(inputCh, doneCh)
	multCh := mult(addCh, doneCh)

	for value := range multCh {
		fmt.Println(value)
	}
}
