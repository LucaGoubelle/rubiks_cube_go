package main

import (
	"fmt"
	"simul/gocube/gocube"
	"simul/gocube/gocube/moves"
	"simul/gocube/gocube/data"
)

func main() {
	// var cube = gocube.GetCube(3)
	var cube = data.ImportFromJSON("res/3x3.json")
	gocube.PrintCube(cube)
	fmt.Printf("\n-------------------------------------------\n\n")
	cube = moves.SimpleMove(cube, "R2")
	gocube.PrintCube(cube)
	data.ExportToJSON(cube, "res/3x3.json")
}
