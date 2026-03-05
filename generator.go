// Generator - паттерн, позволяющий обрабатывать данные одновременно
// отправителю и обработчику. Экономит память, так как данные обрабатываются
// по мере поступления

package Patterns

import "fmt"

// generator - преобразует данные в канал
func generator(a []int) chan int {
	out := make(chan int, len(a))
	go func() {
		defer close(out)
		for _, n := range a {
			out <- n
		}
	}()
	return out
}

// process - потребитель данных
func process(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	channel := generator(slice)

	process(channel)
}
