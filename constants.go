package main

import "fmt"

const myConst int = 78

func main(){
	fmt.Println("Inside constants")

	//typed constant

	const myConst int16 = 100

	fmt.Printf("%v, %T", myConst, myConst)
}