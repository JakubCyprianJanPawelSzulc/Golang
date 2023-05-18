package config

import (
	"flag"
	"runtime"
)

type Config struct {
	Program string
	Maxx    int
	Maxy    int
	Tile    int
}

func Setup(x, y, t int) Config {
	maxx := x
	maxy := y
	tile := t
	runtime.LockOSThread()

	flag.IntVar(&maxx, "x", maxx, "horizontal resolution")
	flag.IntVar(&maxy, "y", maxy, "vertical resolution")
	flag.IntVar(&tile, "t", tile, "pixel size")
	flag.Parse()
	return Config{
		Program: "Przykład",
		Maxx:    maxx,
		Maxy:    maxy,
		Tile:    tile,
	}
}
