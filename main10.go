package main

import "fmt"

func main()  {
	var x int = 10
	var y int = 20
	fmt.Printf("a = %d, y = %d", x, y)

	fmt.Println("swap a & b")

	swap(&x, &y)

	fmt.Printf("a = %d, y = %d", x, y)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}