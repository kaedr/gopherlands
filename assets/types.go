package assets

import (
	"strings"
)

// Asset is used to represent all the things we can load from configs
type Asset interface {
	AssetFullPath()
}

// Block as parsed from our json
type Block struct {
	Name       string
	Pretty     string
	Path       string
	Toughness  int
	Substances []string
	Faces      map[string]string
}

// FullPath for it's fully qualified id
func (b Block) FullPath() string {
	return b.Path + "." + b.Name
}

// TexturePath is the path of the texture directory for this block
func (b Block) TexturePath() string {
	pieces := strings.Split(b.Path, ".")
	// Remove the trailing item
	pieces = pieces[:len(pieces)-1]
	pieces = append([]string{"assets"}, pieces...)
	pieces = append(pieces, "textures")
	return strings.Join(pieces, ".")
}
