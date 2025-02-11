package main

import "fmt"

func main() {

	result := sum(10, 10)

	fmt.Println(result)

	result2 := sum(20, 30)

	fmt.Println(result2)

}

func sum(number1 int, number2 int) int {

	return number1 + number2
}
