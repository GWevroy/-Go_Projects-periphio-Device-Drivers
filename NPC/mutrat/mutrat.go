package mutrat

// Mutant Rat NPC

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Constants
const (
	pathImageIdle = "./NPC/mutrat/assets/images/Idle.png" // Path to player idle animation image
	ratWidth      = 128                                   // Image size for rat
	ratHeight     = 128
)

type Rat struct {
	PosX    float64 // X, Y coordinates for player
	PosY    float64
	imgIdle *ebiten.Image
	health  int8 // Health measured in percentage, or 0 - 100
}

func NewRat(posX, posY float64) *Rat {
	var err error
	rat := Rat{}
	rat.health = 100
	rat.PosX = posX // Assert starting point X and Y player coordinates
	rat.PosY = posY

	rat.imgIdle, _, err = ebitenutil.NewImageFromFile(pathImageIdle)
	if err != nil {
		fmt.Println("Failed to load rat idle image from ", pathImageIdle)
		log.Fatal(err)
	}
	return &rat
}

func (r *Rat) DrawRat(screen *ebiten.Image) {

	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(r.PosX, r.PosY)
	screen.DrawImage(
		r.imgIdle.SubImage(image.Rect(0, 0, ratWidth, ratHeight)).(*ebiten.Image), // Render subimage (withimage.Image type assert back to ebiten.Image)
		&opt)
}
