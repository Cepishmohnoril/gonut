package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 400
	TPS          = 10
	R1           = 2.0
	R2           = 5.0
	K2           = 5.0
	K1           = float64(screenWidth) * K2 * 3.0 / (8.0 * (R1 + R2))
	alphaStep    = 0.01
	betaStep     = 0.01
	twoPi        = math.Pi * 2
)

type Game struct {
	pixels []byte
	gamma  float64
}

func main() {
	g := newGame()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Render donut in Go. (Gonut) This is not a call. =)")
	ebiten.SetTPS(TPS)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func newGame() *Game {
	g := &Game{
		pixels: make([]byte, screenWidth*screenHeight*4),
	}

	return g
}

func (g *Game) renderFrame(A float64, B float64) {
	cosA := math.Cos(A)
	sinA := math.Sin(A)
	cosB := math.Cos(B)
	sinB := math.Sin(B)
	zBuffer := make([][]float64, screenWidth)

	for i := range zBuffer {
		zBuffer[i] = make([]float64, screenHeight)
	}

	for alpha := 0.0; alpha <= twoPi; alpha += alphaStep {
		cosAlpha := math.Cos(alpha)
		sinAlpha := math.Sin(alpha)

		for beta := 0.0; beta <= twoPi; beta += betaStep {
			cosBeta := math.Cos(beta)
			sinBeta := math.Sin(beta)

			x := (R2+R1*cosAlpha)*(cosBeta*cosB+sinA*sinB*sinBeta) - R1*cosA*sinB*sinAlpha
			y := (R2+R1*cosAlpha)*(cosBeta*sinB-cosB*sinA*sinBeta) + R1*cosA*cosB*sinAlpha
			z := cosA*(R2+R1*cosAlpha)*sinBeta + R1*sinA*sinAlpha
			xp := int(math.Round((K1*x/K2 + z)) + 200)
			yp := int(math.Round((K1*y/K2 + z)) + 200)
			L := cosBeta*cosAlpha*sinB - cosA*cosAlpha*sinBeta - sinA*sinAlpha + cosB*(cosA*sinAlpha-cosAlpha*sinA*sinBeta)

			if z > zBuffer[xp][yp] {
				g.setPixel(byte(L*255), xp, yp)
				zBuffer[xp][yp] = z
			}
		}
	}
}

func (g *Game) setPixel(val byte, x int, y int) {
	p := int(4 * (y*screenWidth + x))

	g.pixels[p] = val
	g.pixels[p+1] = val
	g.pixels[p+2] = val
	g.pixels[p+3] = val
}

func (g *Game) Update() error {
	g.gamma += 0.1

	if g.gamma >= twoPi {
		g.gamma = 0
	}

	g.pixels = make([]byte, screenWidth*screenHeight*4)
	g.renderFrame(g.gamma, g.gamma)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
