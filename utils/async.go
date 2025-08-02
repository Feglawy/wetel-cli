package utils

import "sync"

type Result[T any] struct {
	Val T
	Err error
}

func RunTask[T any](
	fn func(string) (T, error),
	arg string,
	wg *sync.WaitGroup,
	ch chan<- Result[T],
) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		val, err := fn(arg)
		ch <- Result[T]{Val: val, Err: err}
	}()
}
