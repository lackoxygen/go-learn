package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	args := os.Args
	for i := 0; i < len(args) ; i ++  {
		fmt.Println(args[i])
	}

	for _, arg := range args[1:] {
		fmt.Printf(arg)
	}

	fmt.Println(strings.Join(args[1:], " "))

	fmt.Println(args[0])
}