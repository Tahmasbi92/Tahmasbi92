package main

import "fmt"

func main() {

	result := sum(10, 10)

	fmt.Println(result)

}

func sum(number1 int, number2 int) int {

	return number1 + number2
}
