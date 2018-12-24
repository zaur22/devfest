package main

import "fmt"

/***************
* каналы для общения между потоками
* буферизированные и небуферизированные
****************/

func main() {
	ch1 := make(chan int)

	go func(ch chan int) {
		val := <-ch
		fmt.Println("Go get from chan", val)
	}(ch1)

	ch1 <- 101

	fmt.Scanln()
}
