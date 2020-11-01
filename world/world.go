package world

import (
	"github.com/g3n/engine/material"
	//"github.com/g3n/engine/math32"
	//"github.com/g3n/engine/texture"
)

const (
	// ChunkSize on x and y axes in Blocks
	ChunkSize = 16
	// WorldHeight on z axis in Blocks
	WorldHeight = 256
)

// Block as parsed from our json
type Block struct {
	name   string
	facing uint8
	faces  map[string]string
}

// TexturedBlock is a block that can be placed in the world
type TexturedBlock struct {
	block    *Block
	textures map[string]*material.Standard
}

// Chunk stores blocks in convenient sections for working with
type Chunk struct {
	X, Y   int
	Blocks [ChunkSize][ChunkSize][WorldHeight]Block
}
