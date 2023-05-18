// graph package, [c] piotao, <piotao@gmail.com>, started about 202304*, @hebron
// for my fellow students and for myself, as a exercise for Golang classes
// 20230430 - refactored and SIMPLIFIED to be very rudimentary (previously has complicated generics and variadics)
// 20230507 - refactoring further, simplification, more go-like idioms, etc. also lots of tests added

// IMPORTANT convention: all functions which have SIMPLE verb only are methods. If function has verb+noun, it returns NEW instance of data
// example: p.Add(...)  -- operates on p internally, p is changed.
//          q:=AddPoint(p,...) -- takes point, creates new point, returns it
// exception: New, works on self but returns something new

package graph

import (
	"log"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// point in general integer format

type Point struct {
	X int
	Y int
}

// defininig new point from different arguments
// p:=new()               - zero point, same as Point{}
// p:=new(point)          - copy
// p:=new(int,int)        - def
// p:=new(int)            - def, like new(int,int)
// p:=new(point,point)    - MIDDLE between two points
// p:=new(point,float)    - scale
// p:=new(float,point)    - scale
// p:=new(point,int,int)  - shift by xy, like p.X+int,p.Y+int
func NewP(a ...any) Point { return Point{}.New(a...) }
func (p Point) New(a ...any) Point {
	switch len(a) {
	case 0:
		return Point{}
	case 1:
		switch v := a[0].(type) {
		case Point:
			return Point{X: v.X, Y: v.Y}
		case int:
			return Point{X: v, Y: v}
		}
	case 2:
		v1 := a[0]
		v2 := a[1]
		switch v1 := v1.(type) {
		case int:
			switch v2 := v2.(type) {
			case int:
				return Point{X: v1, Y: v2}
			}
		case Point:
			switch v2 := v2.(type) {
			case Point:
				return Point{X: (v1.X + v2.X) / 2, Y: (v1.Y + v2.Y) / 2}
			case float64:
				return Point{X: int(float64(v1.X) * v2), Y: int(float64(v1.Y) * v2)}
			}
		case float64:
			switch v2 := v2.(type) {
			case Point:
				return Point{X: int(v1 * float64(v2.X)), Y: int(v1 * float64(v2.Y))}
			}
		}
	case 3:
		v1 := a[0]
		v2 := a[1]
		v3 := a[2]
		switch v1 := v1.(type) {
		case Point:
			switch v2 := v2.(type) {
			case int:
				switch v3 := v3.(type) {
				case int:
					return Point{X: v1.X + v2, Y: v1.Y + v3}
				}
			}
		}
	}
	log.Printf("[w] bad args to Point::New(%#v)\n", a...)
	return Point{}
}

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Point) From(s sdl.Point) {
	p.X = int(s.X)
	p.Y = int(s.Y)
}

// moves point to some other location, internal change, returns itself
// p.Add(int,int)
// p.Add(point)
// p.Add(int,int).Add(p)
func (p *Point) Add(a ...any) *Point {
	switch len(a) {
	case 0:
		return p
	case 1:
		if v, ok := a[0].(Point); ok {
			p.X += v.X
			p.Y += v.Y
		}
	case 2:
		if v1, ok1 := a[0].(int); ok1 {
			if v2, ok2 := a[1].(int); ok2 {
				p.X += v1
				p.Y += v2
			}
		}
	}
	return p
}

// scaling the point
// p.Scale(0.5)
// AddPoint(a,b).Scale(2.0)
func (P Point) Scale(s float32) Point {
	return Point{
		X: int(float32(P.X) * s),
		Y: int(float32(P.Y) * s),
	}
}

// comparing point to the other point
// if p.Eq(q)    == true if both p and q are equal, otherwise false
func (P Point) Eq(a ...any) bool {
	switch len(a) {
	case 0:
		return false
	case 1:
		if v, ok := a[0].(Point); ok {
			return P.X == v.X && P.Y == v.Y
		}
	case 2:
		if v1, ok1 := a[0].(int); ok1 {
			if v2, ok2 := a[1].(int); ok2 {
				return P.X == v1 && P.Y == v2
			}
		}
	}
	log.Printf("[w] bad args to Point::Eq(%#v)\n", a...)
	return false
}

// AddPoint(point,point) - shift by second point
// AddPoint(point,point,point) - shift by two points
// AddPoint(point,int,int) - shift by xy
func AddPoint(p Point, a ...any) Point {
	l := len(a)
	var P Point = p
	if l == 0 {
		return p
	} else if l == 1 {
		v1 := a[0]
		switch v1 := v1.(type) {
		case Point:
			return Point{X: P.X + v1.X, Y: P.Y + v1.Y}
		}
	} else if l == 2 {
		v1, v2 := a[0], a[1]
		switch v1 := v1.(type) {
		case Point:
			switch v2 := v2.(type) {
			case Point:
				return Point{X: p.X + v1.X + v2.X, Y: p.Y + v1.Y + v2.Y}
			}
		case int:
			switch v2 := v2.(type) {
			case int:
				return Point{X: p.X + v1, Y: p.Y + v2}
			}
		}
	}
	log.Printf("[w] bad args to Point::AddPoint(%#v)\n", a...)
	return P
}

// calculates the average of points
// ave := MidPoint(p1,p2)        - middle point between two
// ave := MidPoint(p1,p2,p3)     - middle point betwee three
// ave := MidPoint(a,b,c,d,e,d)  - average on all point
func MidPoint(a ...Point) Point {
	var P Point = Point{}
	for _, v := range a {
		P.X += v.X
		P.Y += v.Y
	}
	l := len(a)
	if l > 0 {
		P.X /= l
		P.Y /= l
	}
	return P
}

func ScalePoint(P Point, s float64) Point {
	return Point{
		X: int(float64(P.X) * s),
		Y: int(float64(P.Y) * s),
	}
}

// draws pixel P using any renderer
func DrawPoint(R *sdl.Renderer, P Point) {
	r, g, b, a, _ := R.GetDrawColor()
	gfx.PixelRGBA(R, int32(P.X), int32(P.Y), r, g, b, a)
}

func (P Point) Draw(G *Graph) {
	gfx.PixelColor(G.Rend, int32(P.X), int32(P.Y), G.F)
}
