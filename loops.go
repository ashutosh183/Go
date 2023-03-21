package main

import "fmt"

func main(){

	for i:=1; i<11; i++{
		fmt.Printf("%v ",i)
	}
	fmt.Println()
	for i:=0; i<5; i++{
		fmt.Println(i)
		if i%2 == 0{
			i /= 2
		}else{
			i = 2 * i + 1
		}
	}

	i := 1
	for{
		fmt.Println("i: ", i)
		i++
		if i == 6{
			break
		}
	}

	for i = 1; i<=5; i++{
		for j := 1; j<=5; j++{
			fmt.Println("I * J: ", i*j)
		}
	}

	s := "ashutosh pandey"

	for idx, val := range s{
		fmt.Println("idx is %v, val is %v", idx, val)
	}
}