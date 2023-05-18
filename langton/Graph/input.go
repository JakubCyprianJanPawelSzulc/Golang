package graph

import "github.com/veandco/go-sdl2/sdl"

type Input interface {
	Listen() bool
	KeyPressed(key sdl.Scancode) bool
	GetCursor() sdl.Point
	GetButtons() uint32
}

type input struct {
	keyboard []uint8
	cursor   sdl.Point
	buttons  uint32
}

func (i *input) Listen() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.MouseButtonEvent:
			if t.State == sdl.PRESSED {
				i.buttons = uint32(t.Button)
			}
			i.cursor.X, i.cursor.Y, i.buttons = sdl.GetMouseState()
		case *sdl.MouseMotionEvent:
			i.cursor.X = t.X
			i.cursor.Y = t.Y
		case *sdl.KeyboardEvent:
			i.keyboard = sdl.GetKeyboardState()
			if i.keyboard[sdl.SCANCODE_ESCAPE] == 1 || i.keyboard[sdl.SCANCODE_Q] == 1 {
				return false
			}
		}
	}
	return true
}

func (i *input) KeyPressed(key sdl.Scancode) bool {
	return i.keyboard[key] == 1
}

func (i *input) GetCursor() sdl.Point {
	return i.cursor
}

func (i *input) GetButtons() uint32 {
	return i.buttons
}

func NewInput() Input {
	return &input{
		keyboard: sdl.GetKeyboardState(),
		cursor:   sdl.Point{},
	}
}

var in Input

func init() {
	in = NewInput()
}

func Listen() bool {
	return in.Listen()
}

func KeyPressed(key sdl.Scancode) bool {
	return in.KeyPressed(key)
}

func GetCursor() sdl.Point {
	return in.GetCursor()
}

func GetButtons() uint32 {
	return in.GetButtons()
}
