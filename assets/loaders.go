package assets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kaedr/gopherlands/utils"
)

// Blocks keeps track of all the blocks that we've loaded
var Blocks = make(map[string]Block)

type blockPack struct {
	Path  string
	Items []Block
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
	return blockPack{filepath, blockList}
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
			blockChan <- loadBlocksFile(blockDir + utils.OPS + file.Name())
		}
	}
}

func compileAssetMap(blockChan chan blockPack, blockMap map[string]Block, waiter *sync.WaitGroup) {
	defer waiter.Done()
	for pack := range blockChan {
		for _, block := range pack.Items {
			blockMap[pack.Path+block.Name] = block
		}
	}
}

// LoadAssets go through all our assets and loads them
func LoadAssets() {
	// Set up waitgroups to track when we're done loading and done mapping
	var loadWaiter, mapWaiter sync.WaitGroup
	// Set up our channels for the loaders to pass assets back across
	blockChan := make(chan blockPack, 50)
	// Spin up our map compilers
	mapWaiter.Add(1)
	go compileAssetMap(blockChan, Blocks, &mapWaiter)

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
	// Now wait for all the map compilers to returns
	mapWaiter.Wait()
}
