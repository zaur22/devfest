package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

/***************
* легче зеленых потоков
* настоящая многопоточность
* не могут работать с данными
***************/

func gorutine(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	var n = 8
	for i := 0; i < 8; i++ {
		fmt.Printf("%v▐%v %v\n",
			strings.Repeat(" ", i),
			strings.Repeat(" ", n-i),
			in,
		)
		runtime.Gosched()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		//нумеруем и запускаем
		go gorutine(i, wg)
	}

	wg.Wait()
}
