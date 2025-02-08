package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	result := counter()

	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())

}
