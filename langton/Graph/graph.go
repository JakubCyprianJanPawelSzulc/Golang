// graph package, [c] piotao, <piotao@gmail.com>, started about 202304*, @hebron
// for my fellow students and for myself, as a exercise for Golang classes

package graph

import (
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Graph struct {
	// public fields
	Title string        // title of the graphical window
	Size  Point         // size of the window - W x H, but Size.X × Size.Y
	Rend  *sdl.Renderer // internal data with rendering system
	Flags Flags         // flags on/off for misc configuration data
	F     sdl.Color     // default foreground color
	B     sdl.Color     // default background color
	// private fields
	win    *sdl.Window // internal data of the window
	fonts  Fonts       // font list loaded
	font   string      // last font used
	colors Colors      // colors defined
	Ticks  int64       // ticks calculated for FPS calculation
	Frames int64       // frames drawn
	FPS    float32     // frames per second
}

//####### SMALL private functions for general graph struct

// guard handler for functions which can return nil or errors
// program will crash, because it can't continue without something
func handleError(err error) {
	if err != nil {
		log.Fatalln("[e] {}: {}", err, err.Error())
	}
}

// creates SDL window and handles errors
func createWin(title string, size Point) *sdl.Window {
	win, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(size.X), int32(size.Y), sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE|sdl.WINDOW_OPENGL|sdl.WINDOW_ALLOW_HIGHDPI)
	handleError(err)
	return win
}

// creates renderer sdl inside window and handles errors
func createRend(W *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(W, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	handleError(err)
	return renderer
}

// opening new window of given size
// Init(string,int,int,bool)
func Init(title string, x, y int, verbose bool) *Graph {
	var G *Graph = &Graph{
		Title:  title,
		Size:   Point{}.New(x, y),
		fonts:  make(Fonts),
		colors: make(Colors),
		Flags:  Flags{}.New(),
		Ticks:  time.Now().UnixMilli(),
	}
	G.Flags["verbose"] = verbose
	G.Flags["init-ttf"] = false
	G.Size = Point{}.New(x, y)
	handleError(sdl.Init(sdl.INIT_EVERYTHING))
	G.Flags.Set("init-all")
	handleError(ttf.Init())
	G.Flags.Set("init-ttf")
	G.win = createWin(G.Title, G.Size)
	G.Rend = createRend(G.win)

	// if !G.loadFonts("fonts") {
	// 	log.Println("[w] Can not load any font! Strange.")
	// } else {
	// 	log.Printf("[i] Successfully loaded %d TTF fonts.", len(G.fonts))
	// }

	if !G.loadColors("Graph/colors/rgb.txt") {
		if !G.loadColors("rgb.txt") {
			log.Println("[w] Can not load colors from Graph/colors/rgb.txt or from local rgb.txt file.")
		}
	} else {
		log.Printf("[i] colors file read in with %d colors\n", len(G.colors))
	}
	G.F = G.AddColor("white", 0xf0f0f0ff) // at least this
	G.B = G.AddColor("black", 0x040404ff) // and this color will be defined
	G.Rend.Clear()
	G.Rend.Present()
	log.Println("[i] Graphics initialized.")
	return G
}

// finalizing SDL calls before closing the program
func (g *Graph) Finish() {
	destroyRend(g)
	destroyWin(g)
	destroyFonts(g)
	sdl.Quit()
	log.Println("[i] Graphics finished.")
}

func destroyWin(G *Graph) { // safely calls destroy on a window kept in main graphics
	if G != nil && G.win != nil {
		G.win.Destroy()
	}
}

func destroyRend(G *Graph) { // safely calls destroy on the internal renderer
	if G != nil && G.Rend != nil {
		G.Rend.Destroy()
	}
}

func destroyFonts(G *Graph) { // safely unloads all fonts and closes them
	if G != nil && G.Flags.IsSet("init-ttf") {
		if G.fonts != nil {
			for _, f := range G.fonts {
				f.Close()
			}
		}
		ttf.Quit()
	}
}

// clears the entire window with background color
func (G *Graph) Clear() {
	G.Rend.Clear()
}

// force window to redraw and present new content
func (G *Graph) Show() {
	G.Frames++
	G.Rend.Present()
}

func (G *Graph) SetWindowTitle(t string) {
	G.win.SetTitle(t)
}

func (G *Graph) CalcFPS(ticks int64) float32 {
	if G.Frames >= 15 && ticks > 0 {
		ticks = time.Now().UnixMilli() - ticks
		G.FPS = 1000.0 * float32(G.Frames) / float32(ticks)
		G.Ticks = time.Now().UnixMilli()
		G.Frames = 0
	}
	return G.FPS
}

//
//
// ##################### UTILITY FUNCTIONS
//

// returns -a if a is negative, works on INT only
func Abs[T int | int8 | int16 | int32 | int64 | float32 | float64](a T) T {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// returns -a if a is negative, works on INT only
func Lim[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](min, v, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func fexists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
