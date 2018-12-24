package main

import "fmt"

/*************************
*  panic это не исключения
*  паника нежелательна
**************************/

func deferTest() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("PAAAANNNIIICC: ", err)
		}
	}()
	defer hello()
	fmt.Println("some text")
	panic("Run Forrest Run!")
	fmt.Print("unreachable code")
	return
}

func main() {
	deferTest()
	return
}

func hello() {
	fmt.Println("hello")
}
