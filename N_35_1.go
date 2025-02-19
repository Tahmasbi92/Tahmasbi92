package main

import "fmt"

func main() {
	mymap := map[int]string{12: "apple", 13: "orange", 14: "banana", 15: "pitch"}

	value, exists := mymap[12]

	if exists {
		fmt.Println("my key was", value)
	} else {
		fmt.Println("my key was not found")
	}

}
