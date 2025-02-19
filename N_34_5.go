package main

import "fmt"

func main() {
	personmap := map[int]string{1234: "Alireza Noorghorbani", 234567: "Mohammad Ranjbar"}

	fmt.Println(personmap[1234])
	fmt.Println(personmap[234567])
	fmt.Println(personmap[3456])
	personmap[5537] = "masoud"
	fmt.Println(personmap[5537])
}
