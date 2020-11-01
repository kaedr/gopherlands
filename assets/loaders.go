package assets

import (
	"encoding/json"
	"log"
	"io/ioutil"

	"github.com/kaedr/gopherlands/world"
)

var Blocks map[string]world.Block

func LoadBlocks() {
    content, err := ioutil.ReadFile("assets/blocks/stones.json")
    if err != nil {
        log.Fatal(err)
    }
	json.Unmarshal([]byte(content), &Blocks)
}
