package main

import "strconv"
import "fmt"

type Person struct {
	name string
	age int
	sex string
	city string
}

func (p Person) getName() string {
	return p.name
}

func (p Person) getAge() int {
	return p.age
}

func (p Person) getSex() string {
	return p.sex
}

func (p Person) getCity() string {
	return p.city
}

func main()  {
	zs := Person{"张三", 20, "男", "广州"}
	ls := Person{"李四", 25, "男", "北京"}
	xf := Person{"小芳", 18, "女", "上海"}

	println("zs：" + zs.getName() + "\t" + strconv.Itoa(zs.getAge()) + "\t" + zs.getCity() + "\t" + zs.getSex())
	println("ls：" + ls.getName() + "\t" + strconv.Itoa(ls.getAge()) + "\t" + ls.getCity() + "\t" + ls.getSex())
	println("xf：" + xf.getName() + "\t" + strconv.Itoa(xf.getAge()) + "\t" + xf.getCity() + "\t" + xf.getSex())

	fmt.Printf("%v", zs)

}
