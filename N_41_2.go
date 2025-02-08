package main

import "fmt"

func main() {

	result := sum(10, 22, 34, 56)
	fmt.Println(result)

}

func sum(nums ...int) int {

	total := 0
	for _, num := range nums {

		// total+=num

		total = total + num

	}

	return total
}
