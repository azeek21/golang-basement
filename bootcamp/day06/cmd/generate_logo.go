package main

import (
	"log"

	"github.com/azeek21/blog/apps/drawer"
)

func main() {
	dw := drawer.NewDrawer()
	err := dw.DrawPngFromFile("./public/logo.dws", "./public/logo.png")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Logo successfully generated!")
}
