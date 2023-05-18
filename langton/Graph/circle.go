//

// graph package, [c] piotao, <piotao@gmail.com>, started about 202304*, @hebron
// for my fellow students and for myself, as a exercise for Golang classes
// CIRCLE related functions

package graph

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Circle
type Circle struct {
	Pos    Point     // position of top-left corner
	Radius int       // width and height of the Circle angle
	F      sdl.Color // foreground color
}

func (c Circle) New(x, y, r int, F RGBA) Circle {
	return Circle{
		Pos:    Point{X: x, Y: y},
		Radius: r,
		F:      F.Color(),
	}
}

// create new Circle  from two points: pos=position, size=height and width, calculates middle pos as O = origin
// example:
// r := NewCircle ()      -- zero-Circle
// r := NewCircle (oldc)  -- copy from other Circle
// r := NewCircle (p,r)   -- if p is point and r is int - simple circle given!
// r := NewCircle (x,y,r) -- position in xy, radius in r
// other combinations: warning and empty Circle  as a result! maybe for further todos
func NewCircle(a ...any) Circle {
	if len(a) == 0 {
		return Circle{}
	} else if len(a) == 1 {
		switch a[0].(type) {
		case Circle:
			return Circle{Pos: a[0].(Circle).Pos, Radius: a[0].(Circle).Radius}
		default:
			fmt.Println("[w] ambigous NewCircle  parameters! - single param which is not Circle!")
		}
	} else if len(a) == 2 {
		// r := NewCircle (p,r)   -- if p is point and r is int - simple circle given!
		switch a[0].(type) {
		case Point:
			switch a[1].(type) {
			case int:
				return Circle{
					Pos:    a[0].(Point),
					Radius: a[1].(int),
				}
			}
		default:
			fmt.Println("[w] ambigous NewCircle  parameters! - expected Point,int")
		}
	} else if len(a) == 3 {
		// r := NewCircle (x,y,r) -- position in xy, radius in r
		switch a[0].(type) {
		case int:
			switch a[1].(type) {
			case int:
				switch a[2].(type) {
				case int:
					return Circle{
						Pos:    Point{X: a[0].(int), Y: a[1].(int)},
						Radius: a[2].(int),
					}
				}
			default:
				fmt.Println("[w] ambigous NewCircle  parameters! - triple params but not int,int,int")
			}
		default:
			fmt.Println("[w] ambigous NewCircle  parameters! - triple params but not int,int,int")
		}
	} else {
		fmt.Println("[w] ambigous NewCircle  parameters - too much of them (needed 0-3 only)")
	}
	return Circle{}
}

func (c *Circle) Draw(G *Graph) {
	gfx.AACircleColor(G.Rend, int32(c.Pos.X), int32(c.Pos.Y), int32(c.Radius), G.F)
}

func (r *Circle) Move(dx, dy int) {
	r.Pos.Add(dx, dy)
}

func (G *Graph) DrawCircle(c Circle) {
	if G != nil {
		gfx.AACircleColor(G.Rend, int32(c.Pos.X), int32(c.Pos.Y), int32(c.Radius), G.F)
	}
}
