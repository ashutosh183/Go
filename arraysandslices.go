package main


import "fmt"

func main(){
	fmt.Println("Inside Arrays")

	//Arrays are used when we have multiple values of same type
	grades := []int{98, 99, 96, 97}

	fmt.Printf("%v, %T \n", grades, grades)


	var students [3]string
	fmt.Printf("%v, %T \n", students, students)
	students[0] = "Ashutosh"

	fmt.Printf("%v \n", students)

	a := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array a before change using pointer: %v \n", a)
	b := &a
	b[1] = 4
	fmt.Printf("Capacity: %v \n", cap(a))
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)


	//Slices

	mySlice := a[:] 
	mySlice1 := a[3:]
	mySlice2 := a[3:5] 
	mySlice3 := a[:5]

	fmt.Printf("%v \n", mySlice)
	fmt.Printf("%v \n", mySlice1)
	fmt.Printf("%v  \n", mySlice2)
	fmt.Printf("%v \n", mySlice3)
}