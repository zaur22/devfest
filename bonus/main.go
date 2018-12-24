package main

import (
	"fmt"
	"sync"
)

func DoIt(num int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("Close %v\n", num)
	for msg := range ch {
		fmt.Printf("%v\t%v\n", num, msg)
	}
}

func main() {
	var ch = make(chan string, 1)
	var wg = &sync.WaitGroup{}
	var input string

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DoIt(i, ch, wg)
	}

	for input != "-" {
		fmt.Scanln(&input)
		ch <- input
	}

	close(ch)
	wg.Wait()
}
