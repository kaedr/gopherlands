package assets

// Asset is used to represent all the things we can load from configs
type Asset interface {
	AssetName()
	AssetPath()
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

// AssetName this block goes by
func (blck Block) AssetName() string {
	return blck.Name
}

// AssetPath where it lives
func (blck Block) AssetPath() string {
	return blck.Path
}

// AssetFullPath for it's fully qualified id
func (blck Block) AssetFullPath() string {
	return blck.Path + "/" + blck.Name
}
