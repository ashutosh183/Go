package main


import "fmt"

func main(){

	number := 50
	guessedNumber := 50

	if guessedNumber >= 1 && guessedNumber <= 100{
		if guessedNumber < number{
			fmt.Println("too low")
		}else if guessedNumber > number{
			fmt.Println("too high")
		}else{
			fmt.Println("Guessed Correctly!! ", number)
		}
	}else{
		fmt.Println("Guess number b/w 1 and 100")
	}

	price := 100

	switch price{
	case 10, 20, 30: fmt.Println(
		"Between 10 and 30")
	case 40,50,60: fmt.Println(
		"Between 40 and 60")
	case 70,80,90, 100: fmt.Println(
		"Between 70 and 100")
	}
}