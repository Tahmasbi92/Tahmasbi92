package main

import "fmt"

func main() {

	defer fmt.Println("this is first") // خطي که defer قبل از آن آمده است آخر از همه اجرا می شود.
	fmt.Println("this is second")
	defer fmt.Println("this is third") // چون در خط اول defer اول آمده است و بعد defer  به خط سوم اضافه شد بنابراین دیرتر از همه اجرا می شود.
}
