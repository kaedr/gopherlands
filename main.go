package main

//"github.com/kaedr/gopherlands/client"
import (
	"log"

	"github.com/kaedr/gopherlands/assets"
)

func main() {
	//client.StartClient()
	assets.LoadAssets()
	log.Println("LoadAssets returned, parsed:")
	log.Println(assets.Blocks)
}
