package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(1024, 512, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	x, y := float32(512), float32(256)

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.Color3f(1.0, 1.0, 1.0)
		gl.Begin(gl.POINTS)
		gl.Vertex2f(x, y)
		gl.End()

		window.SwapBuffers()
		glfw.PollEvents()

		if window.GetKey(glfw.KeyQ) == glfw.Press || window.GetKey(glfw.KeyEscape) == glfw.Press {
			break
		}

		dx, dy := rand.Float32()*2-1, rand.Float32()*2-1
		x += dx
		y += dy

		if x < 0 || x > 1024 || y < 0 || y > 512 {
			break
		}
	}
}
