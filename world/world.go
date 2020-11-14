package world

import (
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

// Chunk stores blocks in convenient sections for working with
type Chunk struct {
	X, Y   int
	Blocks [ChunkSize][ChunkSize][WorldHeight]*assets.Block
}
