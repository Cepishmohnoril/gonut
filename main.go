package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

type Game struct {
	pixels []byte
}

func main() {
	g := newGame()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
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

	betaStep := 0.1
	betaBorder := math.Pi * 2

	R1 := 20.0
	R2 := 5.0

	K2 := 5.0
	K1 := float64(screenWidth) * K2 * 3.0 / (8.0 * (R1 + R2))

	for alpha := 0.0; alpha <= alphaBorder; alpha += alphaStep {
		for beta := 0.0; beta <= betaBorder; beta += betaStep {

			x := (R1 + R2*math.Cos(beta)) * math.Cos(alpha)
			y := (R1 + R2*math.Cos(beta)) * math.Sin(alpha)
			z := R2 * math.Sin(beta)

			xp := math.Round((K1*x/K2 + z)) + 200
			yp := math.Round((K1*y/K2 + z)) + 200

			if yp > 0 && xp < screenHeight && xp > 0 && xp < screenWidth {

				g.setCell(0xff, xp, yp)
			}
		}
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
