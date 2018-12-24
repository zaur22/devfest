package main

import (
	"devfest/5_package/fest"
	"fmt"
)

func main() {
	devFest2018 := fest.NewDevFest(2018)
	devFest2018.AddVisitor("Ashurbekov Zaur")
	fmt.Printf("%+v \n", devFest2018.Visitors)
}
