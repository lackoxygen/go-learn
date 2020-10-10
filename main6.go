package main

import "fmt"

func main()  {
	var tab = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("%v", tab)

	len := len(tab)

	ca := &tab[len - 1]

	ca[1] = 66

	fmt.Printf("\n%v", ca)

	fmt.Printf("\n%v", tab)


	var cate = [3][2] string {
		{"西瓜", "苹果"},
		{"啤酒", "白酒"},
		{"麦当劳", "肯德基"},
	}

	fmt.Printf("\n%v", cate)

	for i := 0; i < 3; i ++ {
		for j := 0; j < 2; j ++  {
			fmt.Printf("\n%d-%d-%s", i, j, cate[i][j])
		}
	}
}
