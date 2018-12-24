package hello

import "fmt"

/*****************************
* тесты это важно
* грамотные тесты - ещё важнее
* нативаная поддержка в языке
* _test в название
******************************/

func SayHelloFor(name string) string {
	return fmt.Sprintf("Hello, %v", name)
}

func fibonacci(number int) []int {
	fib := []int{0, 1}
	for len(fib) <= number {
		last := len(fib) - 1
		fib = append(
			fib,
			fib[last]+fib[last-1],
		)
	}
	return fib
}
