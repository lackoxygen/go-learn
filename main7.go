package main

func main()  {
	var foods = []string{"白切鸡", "猪脚饭", "海南鸡"}

	p(foods)
}

func p(foods []string)  {
	len := len(foods)
	for i := 0; i < len; i ++ {
		println(foods[i])
	}
}