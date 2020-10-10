package main

import "fmt"

type N interface {
	Get() string
	Set(version string)
}

type S struct {
	version string
}

func (s *S) Get() string {
	return s.version
}

func (s *S) Set(version string) {
	s.version = version
}

func main() {
	s := &S{}
	fmt.Printf("%v\n", s)
	s.Set("1.0.1")
	println(s.Get())
}