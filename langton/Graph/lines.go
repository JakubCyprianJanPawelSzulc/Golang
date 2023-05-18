package graph

import (
	"log"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// drawing alpha-blended line with default color, options are varied:
// g.Line(p1,p2)    -- draws line between points
// g.Line(x,p)      -- draws horizontal line as x1,(x2,y)
// g.Line(p,y)      -- draws vertical line as (x,y1),y2
// g.Line(x,y,p)    -- draws line from xy to point p
// g.Line(p,x,y)    -- draws line from point p to xy position
// g.Line(x,y,z,t)  -- draws line between just loose coordinates
// other parameters are not handled, and warning is issued
func (G *Graph) Line(a ...any) {
	if len(a) == 2 {
		// g.Line(p1,p2)    -- draws line between points
		// g.Line(x,p)      -- draws vertical line
		// g.Line(p,y)      -- draws horizontal line
		switch v1 := a[0].(type) {
		case Point:
			switch v2 := a[1].(type) {
			case sdl.Point:
				v1 = a[0].(Point)
				v2 = a[1].(sdl.Point)
				gfx.AALineColor(G.Rend, int32(v1.X), int32(v1.Y), int32(v2.X), int32(v2.Y), G.F)
			case Point: // g.Line(p1,p2)    -- draws line between points
				gfx.AALineColor(G.Rend, int32(a[0].(Point).X), int32(a[0].(Point).Y), int32(a[1].(Point).X), int32(a[1].(Point).Y), G.F)
			case int: // g.Line(p,y)      -- draws horizontal line
				gfx.VlineColor(G.Rend, int32(a[0].(Point).X), int32(a[1].(int)), int32(a[0].(Point).Y), G.F)
			default:
				log.Println("[w] Line called with 2nd param not being Point nor int")
			}
		case int:
			switch a[1].(type) {
			case Point: // g.Line(x,p)      -- draws vertical line
				gfx.HlineColor(G.Rend, int32(a[0].(int)), int32(a[1].(Point).X), int32(a[1].(Point).Y), G.F)
			}
		default:
			log.Println("[w] Line called with 2nd param not being Point nor int")
		}
	} else if len(a) == 3 {
		// g.Line(x,y,p)    -- draws line from xy to point p
		// g.Line(p,x,y)    -- draws line from point p to xy position
		switch a[0].(type) {
		case int:
			switch a[1].(type) {
			case int:
				switch a[2].(type) {
				case Point: // g.Line(x,y,p)    -- draws line from xy to point p
					gfx.AALineColor(G.Rend, int32(a[0].(int)), int32(a[1].(int)), int32(a[2].(Point).X), int32(a[2].(Point).Y), G.F)
				default:
					log.Println("[w] Line called with 3nd params but last is NOT a Point")
				}
			default:
				log.Println("[w] Line called with 3nd params but second is NOT int")
			}
		case Point:
			switch a[1].(type) {
			case int:
				switch a[2].(type) {
				case int: // g.Line(p,x,y)    -- draws line from point p to xy position
					gfx.AALineColor(G.Rend, int32(a[0].(Point).X), int32(a[0].(Point).Y), int32(a[1].(int)), int32(a[2].(int)), G.F)
				default:
					log.Println("[w] Line called with 3nd params but last is NOT an int")
				}
			default:
				log.Println("[w] Line called with 3nd params but second is NOT int")
			}
		default:
			log.Println("[w] Line called with 3nd params but NOT with int,int,Point or Point,int,int")
		}
	} else if len(a) == 4 {
		// g.Line(x,y,z,t)  -- draws line between just loose coordinates
		switch a[0].(type) {
		case int:
			switch a[1].(type) {
			case int:
				switch a[2].(type) {
				case int:
					switch a[3].(type) {
					case int:
						gfx.AALineColor(G.Rend, int32(a[0].(int)), int32(a[1].(int)), int32(a[2].(int)), int32(a[3].(int)), G.F)
					default:
						log.Println("[w] Line called with 4nd params but last is NOT int")
					}
				}
			default:
				log.Println("[w] Line called with 3rd params but it is NOT int")
			}
		default:
			log.Println("[w] Line called with 2nd params but it is NOT int")
		}
	}
}
