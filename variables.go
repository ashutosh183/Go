package main

import "fmt"



func main(){
	//variable declaration
	var i int
	i = 42
	fmt.Println(i)


	//go compiler automatically identifies the type
	j := 56
	fmt.Println(j)

	//format printing
	fmt.Printf("J value is %v and it's type is %T", j, j)
	fmt.Println()

	name := "Ashutosh Pandey"
	fmt.Printf("Name is %s", name)

	fmt.Println()
	price := 87.9
	fmt.Printf("Price is %f \n", price)

	//Boolean -> true or false
	var isEqual bool = true
	fmt.Printf("%v, %T \n", isEqual, isEqual)


	//Integer
	// can have 8, 16, 32, 64
	//go provides us with both signed and unsigned integers
	//can use Big package for larger numbers
	n := 18
	fmt.Printf("%v, %T \n", n, n)


	//Arithmetic operations
	a := 100
	b := 9

	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)


	//Bitwise operations
	fmt.Println("Bitwise operations")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(b >> 3) //divide by 2
	fmt.Println(b << 3) //multiply by 2

	//Logical operations
	fmt.Println("And Operator")
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(true && true)

	fmt.Println("OR Operator")
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println(true || true)

	fmt.Println("NOT Operator")
	fmt.Println(!true)
	fmt.Println(!false)

	//Complex numbers

	fmt.Println("Complex numbers")
	c := 1 + 2i
	d := 3 + 5i

	fmt.Println(c + d)
	fmt.Println(c * d)
	fmt.Println(c - d)
	fmt.Println(c / d)

	fmt.Printf("Real part %v, Imaginary part %v", real(c), imag(c))


	//Strings

	s := "ashutosh is learning go"

	fmt.Println("Strings are immutable")
	fmt.Printf("%v, %T", s, s)
}