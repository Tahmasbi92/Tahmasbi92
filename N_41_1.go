package main

import "fmt"

func main() {
	myjam := compute(2, 2, jame)
	mytafrigh := compute(10, 2, tafrigh)

	fmt.Println(myjam)
	fmt.Println(mytafrigh)

}

type opreation func(int, int) int

func compute(a, b int, op opreation) int {
	return op(a, b)
}

func jame(a, b int) int {

	return a + b
}

func tafrigh(a, b int) int {

	return a - b
}

func zarb(a, b int) int {

	return a * b
}
