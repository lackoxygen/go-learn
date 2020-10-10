package main

func main()  {
	sum := func(x int, y int) int{
		return x + y
	}

	println(sum(1, 2))

	fun := getFun

	fun()()
}

func getFun() func() {
	f := 0
	return func() {
		f ++
		println(f)
	}
}
