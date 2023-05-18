package graph

import (
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Fonts map[string]*ttf.Font

// loads all fonts from 'fonts' dir, as specified, assumes that fonts have the format:
// name-subtype-boldness-variant.ttf
// like: subtype = normal, condensed, narrow, compressed
// boldness = black, bold, regular, light, extralight, thin
// variant = italic (or empty)
func (G *Graph) loadFonts(dir string) bool {
	var success bool = false
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Println("[w] can't load fonts from directory '", dir, "'!", err)
		return false
	}
	count := 0
	for _, n := range []string{"mono", "sans", "serif"} {
		for _, t := range []string{"normal", "narrow", "condensed", "compressed"} {
			for _, v := range []string{"black", "bold", "medium", "regular", "light", "extralight", "thin"} {
				name := "fonts/" + n + "-" + t + "-" + v + ".ttf"
				if fexists(name) {
					G.LoadFont(n+"-"+t+"-"+v, name, 24)
					count += 1
				} else {
					log.Printf("[w] font with key [%s] does not exists here: [%s]", n+"-"+t+"-"+v, name)
				}
			}
		}
	}
	if count > 0 {
		success = true
	}
	return success
}

// loads font to the internal font list, by default, a font.ttf is loaded
// with size of 24 pixels (quite small!)
func (G *Graph) LoadFont(name string, path string, size int) *ttf.Font {
	if G.fonts == nil {
		G.fonts = make(map[string]*ttf.Font)
	}
	if !G.Flags.IsSet("init-ttf") {
		if err := ttf.Init(); err != nil {
			panic(err)
		}
		G.Flags.Set("init-ttf")
	}
	if _, ok := G.fonts[name]; !ok { // nie było jeszcze tego fontu
		font, err := ttf.OpenFont(path, size)
		if err != nil {
			log.Fatal("[w] ", err, path)
			return nil
		} else {
			if G.Flags.IsSet("verbose") {
				log.Printf("[i] found font %s loaded as [%s]:", path, name)
			}
		}
		G.fonts[name] = font
		G.font = name
		return font
	}
	return nil
}

func (G *Graph) CreateTexture(surf *sdl.Surface) *sdl.Texture {
	if surf != nil {
		if tex, err := G.Rend.CreateTextureFromSurface(surf); err == nil {
			return tex
		} else {
			log.Fatal(err)
		}
	}
	log.Fatal("[w] returning nil from CreateTexture :)")
	return nil
}

// rendering final text in renderer in the window, can scale:
// g.RenderText(surf,x,y)   - no scaling
// g.RenderText(suft,x,y,2.0)  -- upscaling
// g.RenderText(suft,x,y,0.5)  -- downscaling
// g.RenderText(suft,x,y,1.5,2.7)  -- non-uniform scaling by x:y, changing the aspect too.
func (G *Graph) RenderText(surf *sdl.Surface, x int, y int, scale ...float32) {
	if surf != nil {
		if tex := G.CreateTexture(surf); tex != nil {
			defer tex.Destroy()
			var dest sdl.Rect
			if len(scale) == 0 { // no scaling
				dest = sdl.Rect{X: int32(x), Y: int32(y), W: surf.W, H: surf.H}
			} else if len(scale) == 1 { // scale uniform
				dest = sdl.Rect{
					X: int32(x),
					Y: int32(y),
					W: int32(float32(surf.W) * scale[0]),
					H: int32(float32(surf.H) * scale[0]),
				}
			} else if len(scale) > 1 { // scale non-uniform
				dest = sdl.Rect{
					X: int32(x),
					Y: int32(y),
					W: int32(float32(surf.W) * scale[0]),
					H: int32(float32(surf.H) * scale[1]),
				}
			}
			G.Rend.Copy(tex, nil, &dest)
		}
	} else {
		log.Fatal("[w] nil texture from", x, y)
	}
}

// returns the font specified by the user BUT also reads in if not loaded yet
// if font is not loaded, fontkey is considered to be name of the font, so if
// verdana.ttf is not loaded, and verdana is fontkey, then verdana.ttf will be loaded
// if it was found inside "fonts/*" subdirectory
func (G *Graph) GetFont(fontkey string) *ttf.Font {
	if font, ok := G.fonts[fontkey]; ok {
		G.font = fontkey
		return font
	} else {
		font := G.LoadFont(fontkey, "fonts/"+fontkey+".ttf", 24)
		G.fonts[fontkey] = font
		G.font = fontkey
		return font
	}
}

// renders 1:1 font with high-quality, returns raw surface
func (G *Graph) PrepareText(text string, fontkey string) *sdl.Surface {
	if font := G.GetFont(fontkey); font != nil {
		if surf, err := font.RenderUTF8Shaded(text, G.F, G.B); err == nil {
			return surf
		} else {
			log.Fatal("[w]", err)
		}
	} else {
		log.Fatal("[w] no font was found for key: ", fontkey)
	}
	return nil
}

// draws on the graph window  a text with selected TTF font 1:1
// showing the actual result require invocation of Present method
func (G *Graph) DrawText(text string, font string, x int, y int, s ...float32) {
	if surf := G.PrepareText(text, font); surf != nil {
		defer surf.Free()
		G.RenderText(surf, x, y, s...)
	}
}

// simpler text drawing, using previously selected font
func (G *Graph) DrawTxy(x, y int, text string) {
	G.DrawText(text, G.font, x, y)
}
