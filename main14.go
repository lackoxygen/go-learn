package main

import (
	"fmt"
	"os"
)

func main()  {
	fp, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", fp)
}