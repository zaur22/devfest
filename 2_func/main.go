package main

import (
	"fmt"
	"log"
)

/********************************************
* 1) Функция может возвращает несколько ответов
* 2) В go нет исключений, обрабатываем ручками
*********************************************/

func funcExample(someMsg string) (string, error) {
	if someMsg == "" {
		return "", fmt.Errorf("Bad value of some message")
	}
	return fmt.Sprintf("you say %v", someMsg), nil
}

func main() {
	str, err := funcExample("hi")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf(str)
}
