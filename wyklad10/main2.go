import(
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"Graph/graph"
	"Graph/config"
)

type Plansza [][]int

func NewPlansza(X, Y int) Plansza {
	var p [][]int = make([][]int, Y)
	for y := 0; y < Y; y++ {
		p[y] = make([]int, X)
	}
	return p
}

func (P *Plansza) Set(x, y, val int) {
	(*P)[y][x] = val
}

func (P *Plansza) Get(x, y, v int) int {
	(*P)[y][x] = v
}

type Ant struct (
	graph.Point
	dir int
	BG graph.RGBA
	FG graph.RGBA
)

func main(){
	cfg := config.Setup(1800,1000)
	tile := 10
	G := graph.Init(cfg.Program, cfg.Maxx)
	mx := cfg.Maxx / tile
	my := cfg.Maxx/tile
	plansza := NewPlansza(mx,my)
	ant:=NewAnt(mx/2,my/2, 0xffffffff, 0xaaaaaaff)

	iter := 0


}

func (A *Ant) Rotate(G *graph.Graph, P *Plansza, tile int){
	Dir :=[...][2]int{
		{1,0},
		{0,-1},
		{-1,0},
		{0,1},
	}
	if P.Get(A.X,A.Y)==0{
		A.dir = (A.dir+1)%4
		P.Set(A.X,A.Y,1)
	}else if P.Get(A.X, A.Y) ==1{
		A.dir = (A.dir -1 +1)%4
		P.Set(A.X, A.Y, 2)
		A.Show(G, tile)

	}else if P.Get(A.X, A.Y) ==2{
		P.Set(A.X,A.Y, 0)
		G.Rend.SetDrawColor(0,0,250,255)
		G.Rend.FillRect(&sdl.Rect{X:int32(A.X*tile),Y: int32(A.Y*tile), W:int32(tile)})
	}else if P.Get(A.X, A.Y)==3{
		A.dir = (A.dir +2)%4
		P.Set(A.X, A.Y, 0)
		G.Rend.SetDrawColor(0,250,250,255)
		G.Rend.FillRect(&sdl.Rect{X:int32(A.X*tile),Y: int32(A.Y*tile), W:int32(tile)})
	}

	
}

func NewAnt(x,y int, c1,c2 graph.RGBA) Ant(
	return 
)