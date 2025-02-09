package main

import "fmt"

func main() {

	result := func(a, b int) int {

		return a + b

	}

	sum := result(2, 2)

	fmt.Println(sum)

}
