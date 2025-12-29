package main

import (
	"fmt"
	"simul/gocube/gocube"
	"simul/gocube/gocube/moves"
)

func main() {
	var cube = gocube.GetCube(3)
	gocube.PrintCube(cube)
	fmt.Printf("\n-------------------------------------------\n\n")
	cube = moves.SimpleMove(cube, "U2")
	gocube.PrintCube(cube)
}
