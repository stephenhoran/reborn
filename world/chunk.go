package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/stephenhoran/reborn/utilities"
	"image/color"
	"log"
	"strconv"
)

type Chunks map[string]*Chunk

type Chunk struct {
	tiles [][]*Tile

	x int
	y int
}

func (c Chunks) NewChunk(x, y int) {
	chunk := Chunk{
		tiles: make([][]*Tile, ChunkSize),
		x:     x,
		y:     y,
	}
	ty := y

	for i := range chunk.tiles {
		tx := x
		tw := make([]*Tile, ChunkSize.Int())
		for t := range tw {
			tile := NewTile()
			tile.SetX(tx)
			tile.SetY(ty)
			tw[t] = tile
			tx += TileSize.Int()
		}
		chunk.tiles[i] = tw
		ty -= TileSize.Int()
	}

	chunkName := "Chunk_" + strconv.Itoa(x) + "_" + strconv.Itoa(y)

	c[chunkName] = &chunk
}

func (c *Chunk) X() int {
	return c.x
}

func (c *Chunk) Y() int {
	return c.y
}

func (c *Chunk) SetX(x int) {
	c.x = x
}

func (c *Chunk) SetY(y int) {
	c.y = y
}

func (c Chunks) FindChunkAtPosition(x, y int) *Chunk {
	chunkX, chunkY := findUnit(x, y, ChunkPixel)
	return c.GetChunk(c.buildChunkName(chunkX, chunkY))
}

func (c Chunks) FindTileAtPosition(x, y int) *Tile {
	chunk := c.FindChunkAtPosition(x, y)
	if chunk == nil {
		return nil
	}

	tileX, tileY := findUnit(x, y, TileSize)
	indexX := utilities.Abs((tileX - chunk.X()) / TileSize.Int())
	indexY := utilities.Abs((tileY - chunk.Y()) / TileSize.Int())

	return chunk.Tile(indexX, indexY)
}

func (c Chunks) buildChunkName(x, y int) string {
	return "Chunk_" + strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func (c Chunks) GetChunk(name string) *Chunk {
	chunk, ok := c[name]
	if !ok {
		log.Println("Cannot find chunk " + name)
	}

	return chunk
}

// Draw is used to draw all of the chunks on the screen for debugging purposed. Currently this will draw every chunk, not
// get what is visible in the viewport.
func (c Chunks) Draw(screen *ebiten.Image, offsetX int, offsetY int) {
	for _, chunk := range c {
		chunk.Draw(screen, offsetX, offsetY)
	}
}

// Draw (chunk) is used to draw the box for an entire chunk.
func (c Chunk) Draw(screen *ebiten.Image, offsetX int, offsetY int) {
	x := c.X()
	y := c.Y()

	// Top Line - Left Line - Bottom Line - Right Line
	ebitenutil.DrawLine(screen, float64(x+offsetX), float64(offsetY-y), float64(x+offsetX+ChunkPixel.Int()), float64(offsetY-y), color.RGBA{R: 48, G: 48, B: 48, A: 255})
	ebitenutil.DrawLine(screen, float64(x+offsetX), float64(offsetY-y), float64(x+offsetX), float64(offsetY-y+ChunkPixel.Int()), color.RGBA{R: 48, G: 48, B: 48, A: 255})
	ebitenutil.DrawLine(screen, float64(x+offsetX), float64(offsetY-y+ChunkPixel.Int()), float64(x+offsetX+ChunkPixel.Int()), float64(offsetY-y+ChunkPixel.Int()), color.RGBA{R: 48, G: 48, B: 48, A: 255})
	ebitenutil.DrawLine(screen, float64(x+offsetX+ChunkPixel.Int()), float64(offsetY-y), float64(x+offsetX+ChunkPixel.Int()), float64(offsetY-y+ChunkPixel.Int()), color.RGBA{R: 48, G: 48, B: 48, A: 255})
}

func (c *Chunk) Tile(x, y int) *Tile {
	return c.tiles[y][x]
}

func (c *Chunk) DrawChunkTiles(screen *ebiten.Image, offsetX int, offsetY int) {
	for _, tileRow := range c.tiles {
		for _, tile := range tileRow {
			tile.Draw(screen, offsetX, offsetY)
		}
	}
}
