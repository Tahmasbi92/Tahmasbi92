package main

import "fmt"

func main() {

	mysum := sum(2, 2)

	fmt.Println(mysum)

	mymin := Minus(5, 2)

	fmt.Println(mymin)

	mydiv := Divided(5, 2)

	fmt.Println(mydiv)

	mymu := Multiplied(3, 3)

	fmt.Println(mymu)

}

//sum

//dvidid

// ..

func sum(a int, b int) int {

	return a + b
}

func Minus(a int, b int) int {

	return a - b
}

func Divided(a int, b int) float32 {

	return float32(a / b)
}

func Multiplied(a int, b int) int {

	return a * b
}
