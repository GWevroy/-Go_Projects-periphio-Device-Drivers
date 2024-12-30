package player

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Constants
const (
	pathImageIdle = "./player/assets/images/Idle.png" // Path to player idle animation image
	defPosX       = 100                               // Default X, y coordinates
	defPosY       = 100
	scale         = 0.75 // Scale character image
)

type Player struct {
	PosX    float64 // X, Y coordinates for player
	PosY    float64
	imgIdle *ebiten.Image
	health  int8 // Health measured in percentage, or 0 - 100
}

func NewPlayer() *Player {
	var err error
	player := Player{}
	player.health = 100
	player.PosX = defPosX // Assert starting point X and Y player coordinates
	player.PosY = defPosY

	player.imgIdle, _, err = ebitenutil.NewImageFromFile(pathImageIdle)
	if err != nil {
		fmt.Println("Failed to load player idle image from ", pathImageIdle)
		log.Fatal(err)
	}
	return &player
}

func (p *Player) DrawPlayer(screen *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Scale(scale, scale)
	opt.GeoM.Translate(p.PosX, p.PosY)
	screen.DrawImage(
		p.imgIdle.SubImage(image.Rect(0, 0, 128, 128)).(*ebiten.Image), // Render subimage (withimage.Image type assert back to ebiten.Image)
		&opt)
}
