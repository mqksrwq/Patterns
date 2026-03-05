// Future/Promise - аналог async/await из JavaScript,
// то есть какая-то функция, от которой ждут результат, но
// при этом ее выполнение не блокирует основной поток программы

package Patterns

import (
	"fmt"
	"time"
)

// Result - структура результата со значением и ошибкой
type Result struct {
	res int
	err error
}

// Promise - функция, которая "обещает" возвратить результат в будущем
func Promise(fn func() (int, error)) chan Result {
	channel := make(chan Result, 1)
	go func() {
		defer close(channel)
		res, err := fn()
		channel <- Result{res, err}
	}()

	return channel
}

func main() {
	someFunc := func() (int, error) {
		time.Sleep(5 * time.Second)
		return 0, fmt.Errorf("error")
	}

	// записываем в переменную канал, который обработал Promise
	future := Promise(someFunc)

	result := <-future
	if result.err != nil {
		fmt.Println(result.err)
	} else {
		fmt.Println(result.res)
	}
}
