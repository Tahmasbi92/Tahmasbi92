package main

import "fmt"

func main() {

	defer fmt.Println("this is first")
	defer fmt.Println("this is second")
	fmt.Println("this is third")
}

// هر خطری اول defer
// بقیه خط ها بیاید آن خطِ کد دیرتر از همه ران می شود.
