package main

import "fmt"

func main() {
	personmap := map[int]string{1234: "Alireza Noorghorbani", 234567: "Mohammad Ranjbar"}

	fmt.Println(personmap[1234])
	fmt.Println(personmap[234567])
}
