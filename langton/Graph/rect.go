//

// graph package, [c] piotao, <piotao@gmail.com>, started about 202304*, @hebron
// for my fellow students and for myself, as a exercise for Golang classes

package graph

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// rectangular rectangle, similar to sdl.Rect, but MINE!
type Rect struct {
	Pos  Point     // position of top-left corner
	Size Point     // width and height of the rectangle
	O    Point     // origin point, can be everywhere, by default it is in the middle
	F    sdl.Color // default color
}

func (R Rect) New(pos, size Point, F RGBA) Rect {
	return Rect{
		Pos:  pos,
		Size: size,
		F:    F.Color(),
		O:    AddPoint(pos, size).Scale(0.5),
	}
}

// create new rect from two points: pos=position, size=height and width, calculates middle pos as O = origin
// example:
// r := NewRect()      -- zero-rect
// r := NewRect(oldr)  -- copy from other rect
// r := NewRect(p,q)   -- if p and q are Point then they are like corners
// r := NewRect(p,w,h) -- position in p, size defined by w and h
// r := NewRect(x,y,s) -- position as integers, s as size
// r := NewRect(x,y,w,h) -- position as integers
// other combinations: warning and empty rect as a result! maybe for further todos.
func NewR(a ...any) Rect { return NewRect(a...) }
func NewRect(a ...any) Rect {
	if len(a) == 0 {
		return Rect{}
	} else if len(a) == 1 {
		switch a[0].(type) {
		case Rect:
			return Rect{Pos: a[0].(Rect).Pos, Size: a[0].(Rect).Size, O: a[0].(Rect).O}
		default:
			fmt.Println("[w] ambigous NewRect parameters! - single param which is not Rect")
		}
	} else if len(a) == 2 {
		switch a[0].(type) {
		case Point:
			switch a[1].(type) {
			case Point:
				return Rect{
					Pos:  Point{X: a[0].(Point).X, Y: a[0].(Point).Y},
					Size: Point{X: Abs(a[1].(Point).X - a[0].(Point).X), Y: Abs(a[1].(Point).Y - a[0].(Point).Y)},
					O:    Point{X: (a[0].(Point).X + a[1].(Point).X) / 2, Y: (a[0].(Point).Y + a[1].(Point).Y) / 2},
				}
			}
		default:
			fmt.Println("[w] ambigous NewRect parameters! - single param which is not Rect")
		}
	} else if len(a) == 3 {
		// r := NewRect(p,w,h) -- position in p, size defined by w and h
		switch a[0].(type) {
		case Point:
			switch a[1].(type) {
			case int:
				switch a[2].(type) {
				case int:
					return Rect{
						Pos:  a[0].(Point),
						Size: Point{X: a[1].(int), Y: a[2].(int)},
						O:    Point{X: a[1].(int) / 2, Y: a[2].(int) / 2},
					}
				}
			default:
				fmt.Println("[w] ambigous NewRect parameters! - triple params, but not Point,int,int")
			}
		// r := NewRect(x,y,s) -- position as integers, s as size
		case int:
			switch a[1].(type) {
			case int:
				switch a[2].(type) {
				case Point:
					return Rect{
						Pos:  Point{X: a[0].(int), Y: a[1].(int)},
						Size: a[2].(Point),
						O:    Point{X: a[2].(Point).X / 2, Y: a[2].(Point).Y / 2},
					}
				}
			default:
				fmt.Println("[w] ambigous NewRect parameters! - triple params but not int,int,Point")
			}
		default:
			fmt.Println("[w] ambigous NewRect parameters! - triple params but not int,int,Point and not Point,int,int")
		}
	} else {
		x := a[0]
		y := a[1]
		w := a[2]
		h := a[3]
		if x, ok1 := x.(int); ok1 {
			if y, ok2 := y.(int); ok2 {
				if w, ok3 := w.(int); ok3 {
					if h, ok4 := h.(int); ok4 {
						return Rect{
							Pos:  Point{X: x, Y: y},
							Size: Point{X: w, Y: h},
							O:    Point{X: x + w/2, Y: y + h/2},
						}
					}
				}
			}
		}
		fmt.Println("[w] ambigous NewRect parameters - too much of them (needed 0-3 only)")
	}
	return Rect{}
}

func (r *Rect) W() int {
	return r.Size.X
}
func (r *Rect) H() int {
	return r.Size.Y
}
func (r *Rect) X() int {
	return r.Pos.X
}
func (r *Rect) Y() int {
	return r.Pos.Y
}
func (r *Rect) Draw(G *Graph) {
	gfx.RectangleColor(G.Rend, int32(r.Pos.X), int32(r.Pos.Y), int32(r.Pos.X+r.Size.X), int32(r.Pos.Y+r.Size.Y), r.F)
}

// moves rect internally, chainging it's OWN data
func (r *Rect) Move(dx, dy int) {
	r.Pos.Add(dx, dy)
	r.O.Add(dx, dy)
}

// returns NEW rect but moved to the new location, the old one is copied/unchanged
func (r Rect) Shift(dx, dy int) *Rect {
	return &Rect{
		Pos:  *r.Pos.Add(dx, dy),
		Size: r.Size,
		O:    *r.O.Add(dx, dy),
	}
}

func (G *Graph) DrawRect(r Rect) {
	gfx.RectangleColor(G.Rend, int32(r.Pos.X), int32(r.Pos.Y), int32(r.Pos.X+r.Size.X), int32(r.Pos.Y+r.Size.Y), r.F)
}
func (G *Graph) DrawRectXYWH(x, y, w, h int, c RGBA) {
	gfx.RectangleColor(G.Rend, int32(x), int32(y), int32(x+w), int32(y+h), c.Color())
}
func (G *Graph) DrawRectXYWHFB(x, y, w, h int, c1, c2 RGBA) {
	r, g, b, a := c2.RGBA()
	gfx.BoxRGBA(G.Rend, int32(x), int32(y), int32(x+w), int32(y+h), r, g, b, a)
	r, g, b, a = c1.RGBA()
	gfx.RectangleRGBA(G.Rend, int32(x), int32(y), int32(x+w), int32(y+h), r, g, b, a)
}

func (G *Graph) DrawRectFB(R Rect, c1, c2 RGBA) {
	r, g, b, a := c2.RGBA()
	gfx.BoxRGBA(G.Rend, int32(R.Pos.X), int32(R.Pos.Y), int32(R.Pos.X+R.Size.X), int32(R.Pos.Y+R.Size.Y), r, g, b, a)
	r, g, b, a = c1.RGBA()
	gfx.RectangleRGBA(G.Rend, int32(R.Pos.X), int32(R.Pos.Y), int32(R.Pos.X+R.Size.X), int32(R.Pos.Y+R.Size.Y), r, g, b, a)
}

func (G *Graph) DrawRectRoundedFB(R Rect, c1, c2 RGBA, rad int) {
	r, g, b, a := c2.RGBA()
	gfx.RoundedBoxRGBA(G.Rend, int32(R.Pos.X), int32(R.Pos.Y), int32(R.Pos.X+R.Size.X), int32(R.Pos.Y+R.Size.Y), int32(rad), r, g, b, a)
	r, g, b, a = c1.RGBA()
	gfx.RoundedRectangleRGBA(G.Rend, int32(R.Pos.X), int32(R.Pos.Y), int32(R.Pos.X+R.Size.X), int32(R.Pos.Y+R.Size.Y), int32(rad), r, g, b, a)
}

// draws filled rectangle, usage:
// DrawBox() - nothing
// DrawBox(rect)
// DrawBox(point,point)
// DrawBox(point,int,int)
// DrawBox(int,int,point)
// DrawBox(int,int,int,int)
func (G *Graph) DrawBox(a ...any) {
	if G != nil {
		if len(a) > 0 {
			switch r := a[0].(type) {
			case Rect:
				gfx.BoxColor(G.Rend, int32(r.Pos.X), int32(r.Pos.Y), int32(r.Pos.X+r.Size.X), int32(r.Pos.Y+r.Size.Y), G.F)
			case sdl.Rect:
				gfx.BoxColor(G.Rend, int32(r.X), int32(r.Y), int32(r.X+r.W), int32(r.Y+r.H), G.F)
			}
		} else if len(a) == 2 {
			switch p := a[0].(type) {
			case Point:
				switch q := a[1].(type) {
				case Point:
					gfx.BoxColor(G.Rend, int32(p.X), int32(p.Y), int32(q.X), int32(q.Y), G.F)
				case sdl.Point:
					gfx.BoxColor(G.Rend, int32(p.X), int32(p.Y), int32(q.X), int32(q.Y), G.F)
				}
			}
		} else if len(a) == 3 {
			switch v := a[0].(type) {
			case Point:
				switch v2 := a[1].(type) {
				case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
					x := int32(toInt64(v2))
					switch v3 := a[2].(type) {
					case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
						y := int32(toInt64(v3))
						gfx.BoxColor(G.Rend, int32(v.X), int32(v.Y), int32(x), int32(y), G.F)
					}
				}
			case int:
				switch w := a[1].(type) {
				case int:
					switch p := a[2].(type) {
					case Point:
						gfx.BoxColor(G.Rend, int32(v), int32(w), int32(p.X), int32(p.Y), G.F)
					}
				}
			}
		} else if len(a) == 4 { // four ints perhaps?
			a, b, c, d := a[0].(int), a[1].(int), a[2].(int), a[3].(int)
			gfx.BoxColor(G.Rend, int32(a), int32(b), int32(c), int32(d), G.F)
			gfx.RectangleColor(G.Rend, int32(a), int32(b), int32(c), int32(d), G.F)
		}
	}
}

func toInt64(n any) int64 {
	switch v := n.(type) {
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	}
	return n.(int64)
}
