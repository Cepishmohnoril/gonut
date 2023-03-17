package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 100
	screenHeight = 100
)

type Game struct {
	pixels []byte
}

func main() {
	g := newGame()

	ebiten.SetWindowSize(screenWidth*8, screenHeight*8)
	ebiten.SetWindowTitle("Render donut in Go. (Gonut) This is not a call. =)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func newGame() *Game {
	g := &Game{
		pixels: make([]byte, screenWidth*screenHeight*4),
	}

	alphaStep := 0.05
	alphaBorder := math.Pi * 2

	for alpha := 0.0; alpha <= alphaBorder; alpha += alphaStep {

		x := math.Round(10*math.Cos(alpha)) + 50
		y := math.Round(10*math.Sin(alpha)) + 50

		g.setCell(0xff, x, y)
	}

	return g
}

func (g *Game) setCell(val byte, x float64, y float64) {
	p := int(4 * (y*screenWidth + x))

	g.pixels[p] = val
	g.pixels[p+1] = val
	g.pixels[p+2] = val
	g.pixels[p+3] = val
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
