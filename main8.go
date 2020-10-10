package main

import "fmt"

const MAX int = 5

func main()  {
	a := []int{1, 2, 3, 4, 5}
	var i int
	for i = 0; i < MAX; i ++{
		fmt.Printf("a[%d] = %d \n", i, a[i])
	}

	var ptr [MAX] *int

	for i = 0; i < MAX; i ++ {
		ptr[i] = &a[i]
	}

	fmt.Printf("%v", ptr)

	for i = 0; i < MAX ; i ++ {
		*ptr[i] = MAX * i
		fmt.Printf("ptr[%d] = %d \n", i, *ptr[i])
	}

}