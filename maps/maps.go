package maps

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

// Map package unmarshals data from the tile map JSON file and configures the map

const (
	filepathTiled = "./maps/assets/Sewer Tileset/TiledMap Editor/TiledMap Editor Example - Sewers.tmx"
)

var gameMap *tiled.Map

func NewTilemap() {
	var err error
	//var tilemap tilemap

	// Parse .tmx file.
	gameMap, err = tiled.LoadFile(filepathTiled)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
}

// Draw routine is called every screen frame update
func Draw(screen *ebiten.Image) {
	// You can also render the map to an in-memory image for direct
	// use with the default Renderer, or by making your own.
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Render just layer 0 to the Renderer.
	err = renderer.RenderLayer(6)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Get a reference to the Renderer's output, an image.NRGBA struct.
	img := renderer.Result
	ebitenImg := ebiten.NewImageFromImage(img)
	opt := ebiten.DrawImageOptions{}
	screen.DrawImage(ebitenImg, &opt)
}
