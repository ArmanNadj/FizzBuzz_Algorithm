package main

import (
	"fmt"
	"math"
)

func main() {
	//using := we can inherently declare the type of our variables, in this case float
	x := 0.0

	//infinite for/while loop
	//I chose values with decimals > .0 to be the breaking point because
	//the domain of FizzBuzz is all possible integers, thus floats do not exist
	//in the valid domain of this algorithm
	for {

		fmt.Print("Enter a value (To exit, enter any value with a decimal > .0): ")
		fmt.Scanln(&x)
		if math.Mod(x, 1) != 0 {
			fmt.Println(x, " :: Not an integer value. Goodbye!")
			break
		}
		var fb = fizzBuzz(x)
		fmt.Println(x, " :: ", fb)
	}
}

func fizzBuzz(num float64) string {
	//to increase efficiency of this algorithm,
	//we have our most specific case at the top of the if chain
	if math.Mod(num, 5) == 0 {
		if math.Mod(num, 3) == 0 {
			return "FizzBuzz"
		}
		return "Buzz"
	}
	if math.Mod(num, 3) == 0 {
		return "Fizz"
	}
	return "neither Fizz nor Buzz"
}
