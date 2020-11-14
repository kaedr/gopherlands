package assets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/kaedr/gopherlands/utils"
)

// BlockDict keeps track of all the blocks that we've loaded
var BlockDict = make(map[string]Block)

// TexDict keeps track of all the Texutures that we've loaded
var TexDict = make(map[string]*material.Standard)

type blockPack struct {
	Path  string
	Items []Block
}

type texPack struct {
	Path    string
	Texture *material.Standard
}

// LoadBlocksFile takes a filepath and returns a list of blocks loaded from it
func loadBlocksFile(filepath string) blockPack {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("Opening ", filepath, " got error:")
		log.Println(err)
	}
	var blockList []Block
	err = json.Unmarshal([]byte(content), &blockList)
	if err != nil {
		log.Println("Parsing ", filepath, " got error:")
		log.Println(err)
	}
	assetPath := strings.ReplaceAll(filepath, "json", "")
	assetPath = strings.ReplaceAll(assetPath, utils.Sep, ".")
	return blockPack{assetPath, blockList}
}

// LoadBlocks loads all the blockfiles into our map
func loadBlocks(blockDir string, blockChan chan blockPack, waiter *sync.WaitGroup) {
	defer waiter.Done()
	fileInfo, err := ioutil.ReadDir(blockDir)
	if err != nil {
		log.Println("Reading ", blockDir, " got error:")
		log.Println(err)
	}
	for _, file := range fileInfo {
		if strings.HasSuffix(file.Name(), ".json") {
			blockChan <- loadBlocksFile(blockDir + utils.Sep + file.Name())
		}
	}
}

func compileBlockMap(blockChan chan blockPack, blockMap map[string]Block, waiter *sync.WaitGroup) {
	defer waiter.Done()
	for pack := range blockChan {
		for _, block := range pack.Items {
			blockMap[pack.Path+block.Name] = block
		}
	}
}

func loadTextureFile(filepath string, texChan chan texPack, waiter *sync.WaitGroup) {
	defer waiter.Done()
	tex, err := texture.NewTexture2DFromImage(filepath)
	if err != nil {
		log.Println("Error loading texture file ", filepath)
		log.Println(err)
	}
	mat := material.NewStandard(math32.NewColor("white"))
	mat.AddTexture(tex)
	texChan <- texPack{strings.ReplaceAll(filepath, utils.Sep, "."), mat}
}

func loadTextures(texDir string, texChan chan texPack, waiter *sync.WaitGroup) {
	defer waiter.Done()
	fileInfo, err := ioutil.ReadDir(texDir)
	if err != nil {
		log.Println("Reading ", texDir, " got error:")
		log.Println(err)
	}
	for _, file := range fileInfo {
		if strings.HasSuffix(file.Name(), ".png") {
			texPath := texDir + utils.Sep + file.Name()
			waiter.Add(1)
			go loadTextureFile(texPath, texChan, waiter)
		}
	}
}

func compileTextureMap(texChan chan texPack, texMap map[string]*material.Standard,
	waiter *sync.WaitGroup) {
	defer waiter.Done()
	for pack := range texChan {
		texMap[pack.Path] = pack.Texture
	}
}

// LoadAssets go through all our assets and loads them
func LoadAssets() {
	// Set up waitgroups to track when we're done loading and done mapping
	var loadWaiter, mapWaiter sync.WaitGroup
	// Set up our channels for the loaders to pass assets back across
	blockChan := make(chan blockPack, 50)
	texChan := make(chan texPack, 50)
	// Spin up our map compilers
	mapWaiter.Add(2)
	go compileBlockMap(blockChan, BlockDict, &mapWaiter)
	go compileTextureMap(texChan, TexDict, &mapWaiter)

	// Walk our assets directory, delegate to appropriate subloaders
	err := filepath.Walk("assets",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				// Figure out how to dispatch this directory
				switch dirName := info.Name(); dirName {
				// handle blocks directories
				case "blocks":
					loadWaiter.Add(1)
					go loadBlocks(path, blockChan, &loadWaiter)
				case "textures":
					loadWaiter.Add(1)
					go loadTextures(path, texChan, &loadWaiter)
				}
			}
			return nil
		})
	if err != nil {
		log.Println("Asset walker threw error:")
		log.Println(err)
	}

	// Wait for all of our loaders to return
	loadWaiter.Wait()
	close(blockChan)
	close(texChan)
	// Now wait for all the map compilers to returns
	mapWaiter.Wait()
}
