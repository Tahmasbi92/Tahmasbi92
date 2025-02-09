package main

import "fmt"

func main() {
	r, b, m := divide(10, 2)

	fmt.Println(r, b, m)
}

func divide(a int, b int) (int, int, string) {
	result := a / b
	baghimande := a % b
	var messege string
	if baghimande != 0 {
		messege = "baghimande darad bakhsh pazir nist"
	} else {
		messege = "baghimande nadarad bakhsh pazir hast!"

	}

	return result, baghimande, messege
}
