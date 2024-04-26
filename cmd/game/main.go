package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	DefaultWindowWidth  = 340
	DefaultWindowHeight = 240
)

func main() {
	g := &game{
		windowWidth:  DefaultWindowWidth,
		windowHeight: DefaultWindowWidth,
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Game")

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type game struct {
	windowWidth  int
	windowHeight int
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world!")
}

func (g *game) Layout(_, _ int) (int, int) {
	return g.windowWidth, g.windowHeight
}
