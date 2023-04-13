package main

import (
	"flag"
	"fmt"
	"math"
)

func t(a, b, c float64) (x [2]float64, e bool) {
	d := b*b - 4*a*c
	if d < 0 {
		return
	}
	e = true
	x[0] = (-b + math.Sqrt(d)) / (2 * a)
	x[1] = (-b - math.Sqrt(d)) / (2 * a)
	return
}
func main() {
	a, b, c := flag.Float64("a", 0, ""), flag.Float64("b", 0, ""), flag.Float64("c", 0, "")
	flag.Parse()
	fmt.Print(t(*a, *b, *c))
}
