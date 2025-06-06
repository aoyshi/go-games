package main

import (
	"embed"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

var GhostSprite = mustLoadImage("assets/character-ghost.png")

func mustLoadImage(imgPath string) *ebiten.Image {
	file, err := assets.Open(imgPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	position Vector
}

func (g *Game) Update() error {
	speed := 3.0

	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	// Check for diagonal movement and normalize so player isn't faster when moving diagonally
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	g.position.X += delta.X
	g.position.Y += delta.Y

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.position.X, g.position.Y)

	screen.DrawImage(GhostSprite, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		position: Vector{X: 300, Y: 200},
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
