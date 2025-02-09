package main

import "fmt"

func main() {
	r, b := divide(10, 3)

	fmt.Println(r, b)
}

func divide(a int, b int) (int, int) {
	result := a / b
	baghimande := a % b

	return result, baghimande
}
