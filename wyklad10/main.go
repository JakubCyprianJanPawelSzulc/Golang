func main(){
	cfg := config.Setup(1800,1000)
	tile := 10
	G := graph.Init(cfg.Program, cfg.Maxx, cfg.Maxy, false)
	G.LoadFont("narrow", "Graph/fonts/sans-condensed-light.ttf", 24)
	G.DrawNext("Klasyczna mrówka langtona. "+fmt.Sprintf("Box:[%d], tile"), "narrow", 10 10)
	mx := cfg.Maxx / tile
	my := cfg.Maxy / tile
	plansza := NewPlansza(mx,my)
	ant := NewAnt(mx/2,my/2, 0xfffffff, 0xaaaaaaff)
	iter := 0
	for graph/Listen(){
		iter++
		ant.Rotate(G, &plansza, tile)
		if iter%5==0{
			G.DrawText(fmt.Sprintf("Iteracja: %d", iter), "narrow", 10, 10, 0x000000ff)
		}
	}
}