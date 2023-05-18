// przykład nakładki upraszczającej na SDL2 (niedokończony)
// [c]piotao, <piotao@gmail.com> 4go, @hebron, 202304-202305+

// program ilustruje symulację Mrówki Langotona
// ale nie wiem (bo nie sprawdziłem) czy jest to IDEALNIE taka
// sama mrówka jak na wiki czy gdzieś tam w internecie.
// pobieżne obserwacje układu autostrady sugerują, że jest to
// wzór podobny, tylko o odwróconych kolorach. Zatem MRÓWKA jest
// poprawna, kolory to sprawa umowna.

// aby uruchomić cały program, warto rozpakować go, i napisać
// go get github.com/veandco/go-sdl2/sdl
// go mod tidy
// go build

// SDL musi być zainstalowany w systemie niezależnie za pomocą np.
// instalatora pakietów (pod Linuxem) lub jakoś ręcznie w windoze.
// w systemach opartych o DEB można to zrobić pisząc polecenia podobne do:
// $ sudo apt update
// $ sudo apt install nazwa-pakietu
// moja lista pakietów SDL zainstalowanych w systemie:
// libsdl2-2.0-0
// libsdl2-dev
// libsdl2-gfx-1.0-0
// libsdl2-gfx-dev
// libsdl2-image-2.0-0
// libsdl2-image-dev
// libsdl2-mixer-2.0-0
// libsdl2-mixer-dev
// libsdl2-net-2.0-0
// libsdl2-net-dev
// libsdl2-ttf-2.0-0
// libsdl2-ttf-dev

package main

import (
	"fmt"
	graph "langton/Graph"
	"langton/config"
)

// Mrówka używa planszy tylko do sprawdzenia jaki jest kolor danego pola
type Plansza [][]int

// robienie nowej planszy do mrówki, która trzyma "kolory" w postaci liczb
func NewPlansza(X, Y int) Plansza {
	var p [][]int = make([][]int, Y)
	for y := 0; y < Y; y++ {
		p[y] = make([]int, X)
	}
	return p
}

// pobranie wartości z planszy ze współrzędnych xy, wynik: liczba
// a := p.Get(12,15)     // w a jest liczba, np. 0
func (P *Plansza) Get(x, y int) int {
	return (*P)[y][x]
}

// ustawienie wartości v w planszy na pozycji xy:
// p.Set(12,15,1)   // ustawi we współrzędnych [15][12] wartość 1
func (P *Plansza) Set(x, y, v int) {
	(*P)[y][x] = v
}

// mrówka jest strukturą przechowującą różne informacje o sobie
type Ant struct {
	graph.Point     // współrzędne xy w postaci punktu
	dir         int // kierunek w którym idzie mrówka: 0→ 1↑ 2← 3↓
	bok         int // długość boku mrówki (żeby wiedziała, jak duży kwadrat rysować)
}

// robienie nowej mrówki - w sumie zwracana jest prosta struktura z danymi
// mrówka := NewAnt(10,12)
func NewAnt(x, y, t int) Ant {
	return Ant{
		Point: graph.Point{
			X: x,
			Y: y,
		},
		dir: 0,
		bok: t,
	}
}

// wykonuje całą robotę z mrówką lokalnie w jakimś miejscu planszy
// mrówka jest obracana (zmienia kierunek), oraz zmienia położenie
// plansza zmienia kolor na swoich polach
// ant.Zasuwaj(G,&plansza,tile)
func (A *Ant) Zasuwaj(G *graph.Graph, P *Plansza) {

	// kierunki chodzenia mrówki w postaci zmian współrzędnych xy
	// chodzenie jest możliwe tylko w 4 kierunkach: 0→ 1↑ 2← 3↓
	Dir := [...][2]int{
		{1, 0},  // right 0 →
		{0, -1}, // up    1 ↑
		{-1, 0}, // left  2 ←
		{0, 1},  // down  3 ↓
	}

	// algorytm sprawdza jaki jest kolor planszy POD mrówką w aktualnej pozycji
	// dlatego ile kolorów - tyle różnych ifów. Kolory w planszy nie są prawdziwe,
	// to tylko zwykłe numery - 0 to pierwszy kolor, 1 to drugi, itp. Same kolory
	// są rysowane bezpośrednio odpowiednimi funkcjami do rysowania prostokątów

	if P.Get(A.X, A.Y) == 0 { // jeżeli kolor planszy P w pozycji mrówki A to 0...
		A.dir = (A.dir + 1) % 4 // zmień kierunek mrówki na bardziej w lewo
		P.Set(A.X, A.Y, 1)      // ustaw następny kolor w planszy
		A.Draw(G, 0xf0f0f0ff)   // narysowanie pola w kolorze białym
		//
	} else if P.Get(A.X, A.Y) == 1 { // jeżeli kolor planszy P na pozycji mrówki A to 1 ...
		A.dir = (A.dir - 1 + 4) % 4 // skręć w prawo
		P.Set(A.X, A.Y, 2)          // ustaw kolor następny po 1, czyli 2
		A.Draw(G, 0xff0000ff)       // narysowanie pola w kolorze niebieskim
		//
	} else if P.Get(A.X, A.Y) == 2 { // dalej jest tak samo. Aż się prosi żeby to przepisać bardziej ogólnie
		// w przypadku koloru "2" mrówka idzie dalej prosto, więc nie zmieniamy kierunku
		P.Set(A.X, A.Y, 3)    // ustaw kolor numer 3
		A.Draw(G, 0x00ff00ff) // pole w kolorze zielonym
		//
	} else if P.Get(A.X, A.Y) == 3 {
		A.dir = (A.dir + 2) % 4 // mrówka zawraca gdy napotka pole w tym kolorze
		P.Set(A.X, A.Y, 0)      // zmiana koloru na pierwszy
		A.Draw(G, 0x000000ff)
	}

	// po zmianach koloru planszy i rysowaniu, czas na obliczenie nowego położenia mrówki
	mx := len((*P)[0])                   // wymiar poziomy planszy: liczba kolumn, współrzędna max x
	my := len(*P)                        // wymiar pionowy planszy: liczba wierszy, współrzędna max y
	x := (Dir[A.dir][0] + A.X + mx) % mx // nowe położenie mrówki A.x
	y := (Dir[A.dir][1] + A.Y + my) % my // nowe położenie mrówki A.y
	A.Set(x, y)                          // ustawienie mrówki w docelowym punkcie
}

// funkcja skojarzona do rysowania mrówki. Używa danych z A, rysuje za pomocą grafiki z G.
func (A *Ant) Draw(G *graph.Graph, kolor uint32) {
	G.DrawRectXYWHFB(A.X*A.bok, A.Y*A.bok, A.bok, A.bok, graph.RGBA(uint32(0)), graph.RGBA(kolor))
}

func main() {
	// standardowo, jak to u mnie: konfiguracja (zajrzyj do modułu config, tam jest prosta analiza opcji)
	cfg := config.Setup(1800, 1000, 10)
	tile := cfg.Tile // rozmiar kafelka "piksela"

	// inicjalizacja grafiki i otwarcie nowego okna
	G := graph.Init(cfg.Program, cfg.Maxx, cfg.Maxy, false)

	// ładowanie kozackich czcionek i rysowanie tekstu
	G.LoadFont("font", "Graph/fonts/sans-normal-medium.ttf", 24)
	G.DrawText("Klasyczna Mrówka Langtona. Koniec programu: ESC lub q.", "font", 10, 10)

	// robimy współrzędne wirtualnej planszy (kratek)
	mx := cfg.Maxx / tile // wirtualna rozdzielczość x - tyle kratek ile się zmieści dla danego rozmiaru tile
	my := cfg.Maxy / tile // y

	// nowa, pusta plansza
	plansza := NewPlansza(mx, my)

	// nowa mrówka na środku
	ant := NewAnt(mx/2, my/2, cfg.Tile)

	iter := 0 // numer iteracji

	// pętla główna w której sprawdzamy stan program i rysujemy zmiany
	for graph.Listen() {

		iter++ // zliczamy ile razy program wykonał główną pętlę

		ant.Zasuwaj(G, &plansza) // obliczamy działania mrówki, zmieniając planszę

		// dla przyspieszenia wyświetlania: rysujemy napis oraz planszę co 7 iteracji
		if iter%7 == 0 {
			G.DrawText(fmt.Sprintf("Iteracja: %d  ", iter), "font", 10, 40)
			G.Show()
		}

	}

	// i tyle.
	G.Finish()
}
