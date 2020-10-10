package main

const(
	a int = 1
	b int = 2
	c int = 3
	d int = 4
)

func main()  {
	println("a + b =", sum(a, b))

	s1,s2 := sum2(a, b)

	println("a + b & a *b", s1, s2)

	p1, p2 := swap(a, b)

	println("p1 p2", p1, p2)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (int, int) {
	r1 := a + b
	r2 := a * b
	return r1, r2
}

