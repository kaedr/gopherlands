package assets

// Asset is used to represent all the things we can load from configs
type Asset interface {
	AssetFullPath()
}

// Block as parsed from our json
type Block struct {
	Name       string
	Path       string
	Toughness  int
	Substances []string
	Faces      map[string]string
}

// FullPath for it's fully qualified id
func (blck Block) FullPath() string {
	return blck.Path + "/" + blck.Name
}
