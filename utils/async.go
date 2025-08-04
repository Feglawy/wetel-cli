package utils

type Result[T any] struct {
	Val T
	Err error
}

func RunTask[T any](
	fn func(string) (T, error),
	arg string,
	ch chan<- Result[T],
) {
	go func() {
		val, err := fn(arg)
		ch <- Result[T]{Val: val, Err: err}
		close(ch)
	}()
}
