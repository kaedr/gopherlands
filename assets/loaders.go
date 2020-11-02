package assets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kaedr/gopherlands/world"
)

// Blocks keeps track of all the blocks that we've loaded
var Blocks map[string]world.Block

// LoadBlocksFile takes a filepath and returns a list of blocks loaded from it
func LoadBlocksFile(filepath string) []world.Block {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
	}
	var blockList []world.Block
	json.Unmarshal([]byte(content), &blockList)
	return blockList
}

// LoadBlocks loads all the blockfiles into our map
func LoadBlocks() {
}

// LoadAssets go through all our assets and loads them
func LoadAssets() {
	err := filepath.Walk("assets",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				switch dirName := info.Name(); dirName {
				case "blocks":
					fmt.Println("Found blocks dir: ", path)
				default:
					fmt.Println("Some other thing")
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
