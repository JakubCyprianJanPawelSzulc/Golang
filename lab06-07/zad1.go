package main

import (
	"flag"
	"fmt"
	"math"
)

type l struct {
	a, b, c float64
}

func t(s *l) (x [2]float64, e bool) {
	d := s.b*s.b - 4*s.a*s.c
	if d < 0 {
		return
	}
	e = true
	x[0] = (-s.b + math.Sqrt(d)) / (2 * s.a)
	x[1] = (-s.b - math.Sqrt(d)) / (2 * s.a)
	return
}
func main() {
	a, b, c := flag.Float64("a", 0, ""), flag.Float64("b", 0, ""), flag.Float64("c", 0, "")
	flag.Parse()
	w := l{*a, *b, *c}
	fmt.Print(t(&w))
}
