package main

func main() {

	masahat, mohit := rectangleDimesions(8, 3)

	println(masahat, mohit)

}

func rectangleDimesions(x int, y int) (area int, perimeter int) {
	area = x * y
	perimeter = 2 * (x + y)

	return area, perimeter
}
