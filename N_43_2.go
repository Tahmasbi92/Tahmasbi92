package main

import "fmt"

func main() {

	defer fmt.Println("this is first") // خطي که defer قبل از آن آمده است آخر از همه اجرا می شود.
	fmt.Println("this is second")
	fmt.Println("this is third")
}
