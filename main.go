package main

import (
	"fmt"
	"log"
	"os"

	input "FDFreader/controls"
	mmap "FDFreader/model"
	fdf "FDFreader/parsers"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("can't open file")
	}
	data, err := fdf.MustLoad(os.Args[1])
	if err != nil {
		log.Fatalf("failed to parse")
	}

	var angleX, angleY, zoom float64
	fmt.Print("pls enter values of angles in radians")
	fmt.Scan(&angleX, &angleY, &zoom)
	mapp := mmap.InitMap(data)
	app := input.InitInput(mapp, angleX, angleY, zoom)

	ebiten.SetWindowTitle("FDF reader")
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
