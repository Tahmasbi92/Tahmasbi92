package main

import "fmt"

func main() {

	sa := map[int]string{1: "alireza", 2: "mohammad", 3: "maraym", 4: "yashar", 5: "mostafa", 6: "masoud"}

	for key, value := range sa {
		fmt.Printf("key: %d, value: %s", key, value)
	}

}
