package main

import (
	"log"
	"time"

	"github.com/kaedr/gopherlands/assets"
	//"github.com/kaedr/gopherlands/client"
)

func main() {
	start := time.Now()
	assets.LoadAssets()
	elapsed := time.Since(start)
	log.Println("LoadAssets returned, parsed:")
	log.Println(assets.BlockDict)
	log.Println(assets.TexDict)
	log.Println("In ", elapsed)
	//client.StartClient()
}
