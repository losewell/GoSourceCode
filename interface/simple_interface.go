package main

import "fmt"

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	x int
}

func (s Simple) Get() int {
	return s.x
}

func (s Simple) Set(i int) {
	fmt.Println(i)
	s.x = i
	fmt.Println(s.x)
	return
}

func main() {
	sim := Simple{x:5}
	var s Simpler
	s = sim
	fmt.Println(s.Get())
	s.Set(100)
	fmt.Println(sim.Get())
}
