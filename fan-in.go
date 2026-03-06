// Fan-In - позволяет слить множество каналов в один

package Patterns

import "sync"

func FanIn(doneCh chan struct{}, channels ...chan int) chan int {
	resultCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		ints := ch
		go func() {
			defer wg.Done()
			for num := range ints {
				select {
				case <-doneCh:
					return
				case resultCh <- num:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(doneCh)

	}()

	return resultCh
}
