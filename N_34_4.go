package main

import "fmt"

func main() {

	personmap2 := map[int]string{}
	personmap2[234567] = "mostafa"
	personmap2[456789] = "hassan"

	fmt.Println(personmap2[124567])
}
