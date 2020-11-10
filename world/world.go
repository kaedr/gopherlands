package world

import (
	"github.com/g3n/engine/material"
	"github.com/kaedr/gopherlands/assets"
	//"github.com/g3n/engine/math32"
	//"github.com/g3n/engine/texture"
)

const (
	// ChunkSize on x and y axes in Blocks
	ChunkSize = 16
	// WorldHeight on z axis in Blocks
	WorldHeight = 256
)

// TexturedBlock is a block that can be placed in the world
type TexturedBlock struct {
	block    *assets.Block
	textures map[string]*material.Standard
}

// Chunk stores blocks in convenient sections for working with
type Chunk struct {
	X, Y   int
	Blocks [ChunkSize][ChunkSize][WorldHeight]assets.Block
}
