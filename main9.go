package main

import "fmt"

func main()  {
	var a int = 10
	var ptr *int
	var pptr **int
	var ppptr ***int

	ptr = &a

	pptr = &ptr

	ppptr = &pptr

	**pptr = 1000

	***ppptr = 10000


	fmt.Println(a)
	fmt.Println(*ptr)
	fmt.Println(**pptr)
	fmt.Println(***ppptr)

}