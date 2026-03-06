// Semaphore - паттерн, позволяющий ограничить число одновременно работающих
// потоков

package Patterns

import "sync"

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(max int) *Semaphore {
	return &Semaphore{make(chan struct{}, max)}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}

func main() {
	var wg sync.WaitGroup

	s := NewSemaphore(3)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer s.Release()

			s.Acquire()

		}()
	}
	wg.Wait()
}
