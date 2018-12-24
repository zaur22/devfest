package main

import "fmt"

/*******************************
* В go нет классов
* Зато есть стркутуры с методами
* Конструкторы вручную
********************************/

func NewDevFest(year int) DevFest {
	return DevFest{
		Festival: Festival{
			Topics: make(map[string]string),
		},
	}
}

type Festival struct {
	Topics   map[string]string
	Visitors []string
}

type DevFest struct {
	Festival
	Year int
}

func (f *Festival) AddVisitor(name string) {
	f.Visitors = append(f.Visitors, name)
}

func main() {
	devFest2018 := NewDevFest(2018)
	devFest2018.AddVisitor("Ashurbekov Zaur")
	fmt.Printf("%+v \n", devFest2018.Visitors)
}
