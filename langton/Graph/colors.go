//

// colors module, all colors related functions
// [c] piotao, 20230427, for Graph module: easier SDL2

// uwaga: dużo kodu z tego modułu jest w ogóle NIEPOTRZEBNE!!!

package graph

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// flags used to pack/unpack colors from uint32 byte format
const (
	RED   uint32 = 0xFF000000
	GREEN uint32 = 0x00FF0000
	BLUE  uint32 = 0x0000FF00
	ALPHA uint32 = 0x000000FF
)

// compressed type of the color
type RGBA uint32

// conversion to read struct color from packed 4-byte integer
func (c RGBA) Color() sdl.Color {
	return sdl.Color{
		R: uint8((c >> 24) & 0xFF),
		G: uint8((c >> 16) & 0xFF),
		B: uint8((c >> 8) & 0xFF),
		A: uint8((c >> 0) & 0xFF),
	}
}
func (c RGBA) RGBA() (r uint8, g uint8, b uint8, a uint8) {
	r = uint8((c >> 24) & 0xFF)
	g = uint8((c >> 16) & 0xFF)
	b = uint8((c >> 8) & 0xFF)
	a = uint8((c >> 0) & 0xFF)
	return r, g, b, a
}
func (c *RGBA) From(r uint8, g uint8, b uint8, a uint8) {
	*c = RGBA(uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a))
}
func rgbaFrom(r uint8, g uint8, b uint8, a uint8) RGBA {
	return RGBA(uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a))
}

// color list in a map format: name+rgba value, most compressed format
type Colors map[string]RGBA

// single color in struct mode: fields are components of red, green and blue channels, alpha is transparency
type Color struct {
	R, G, B, A uint8
}

// conversion to binary u32 format, from R,G,B,A to 0xrrggbbaa
func (c Color) RGBA() RGBA {
	return RGBA(((uint32(c.R) << 24) & RED) | ((uint32(c.G) << 16) & GREEN) | ((uint32(c.B) << 8) & BLUE) | uint32(c.A))
}

// returns random color
func RandomColor() RGBA {
	return RGBA(rand.Uint64() | 255)
}

func (G *Graph) SetF(c RGBA) {
	G.Rend.SetDrawColor(c.RGBA())
}

// adds colors to the list, AND:
// if typ == nil, do nothing more:     SetColor(nil,"czarny",0,0,0,255)
// if *typ == true, changes fg color   SetColor(FG_COLOR,"czarny",0,0,0,255)
// if *typ == false, changes bg color  SetColor(BG_COLOR,"czarny",0,0,0,255)
func (G *Graph) AddColor(name string, col RGBA) sdl.Color {
	if G.colors == nil {
		G.colors = make(Colors)
	}
	G.colors[name] = col
	return col.Color()
}

func (G *Graph) GetColor(name string) RGBA {
	if c, ok := G.colors[name]; ok {
		return c
	}
	log.Println("[w] GetColor called with '" + name + "' color, which is not found.")
	return 0
}

func (G *Graph) SetForeground(name string, col ...RGBA) {
	if len(col) > 0 {
		G.colors[name] = col[0]
	}
}

func toUint8(s string) uint8 {
	if v, err := strconv.Atoi(s); err == nil {
		return uint8(v)
	} else {
		return 0
	}
}

func (G *Graph) loadColors(path string) bool {
	var success bool = false
	if fexists(path) {
		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("[w] can't read [%s] as RGB definitions file. Reason: %v", path, err.Error())
			return false
		}
		lines := strings.Split(string(content), "\n")
		re := regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+(\d+)\s+(.*)\s*$`)
		for i, line := range lines {
			if len(line) == 0 || line[0] == '!' {
				continue
			}
			match := re.FindStringSubmatch(line)
			if match != nil {
				r := toUint8(match[1])
				g := toUint8(match[2])
				b := toUint8(match[3])
				n := match[4]
				G.colors[n] = rgbaFrom(r, g, b, 255)
				success = true
			} else {
				log.Println("[w] wrong color definition found at line", i, "in colors file:", path)
			}
		}
	} else {
		log.Printf("[w] can't find file [%s] with RGB definitions.", path)
		return false
	}
	return success
}
