package chest

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
	pathImageIdle = "./objects/chest/assets/images/Rustic chest.png" // Path to player idle animation image
	chestWidth    = 64                                               // Image size for rat
	chestHeight   = 64
)

type Chest struct {
	PosX    float64 // X, Y coordinates for player
	PosY    float64
	imgIdle *ebiten.Image
}

// NewChest instantiates a new chest.
// Must be initialized by main game prior to use in game loop/updates/draws
func NewChest(posX, posY float64) *Chest {
	var err error
	chest := Chest{}
	chest.PosX = posX // Assert starting point X and Y player coordinates
	chest.PosY = posY

	chest.imgIdle, _, err = ebitenutil.NewImageFromFile(pathImageIdle)
	if err != nil {
		fmt.Println("Failed to load rustic chest image from ", pathImageIdle)
		log.Fatal(err)
	}
	return &chest
}

// DrawChest is called by Ebiten Draw every game frame
func (c *Chest) DrawChest(screen *ebiten.Image) {

	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(c.PosX, c.PosY)
	screen.DrawImage(
		c.imgIdle.SubImage(image.Rect(0, 0, chestWidth, chestHeight)).(*ebiten.Image), // Render subimage (withimage.Image type assert back to ebiten.Image)
		&opt)
}
