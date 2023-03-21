package main

import (
	"fmt"
)

func main(){
	// & -> address of
	// * -> value at


	a := 42
	b := a

	fmt.Println(a, b)

	fmt.Println(&a, &b)

	var c int = 18
	var d *int = &c


	fmt.Println(&c, d)
}