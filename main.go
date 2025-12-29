package main

import (
	"simul/gocube/gocube"
)

func main() {
	var cube = gocube.GetCube(3)
	gocube.PrintCube(cube)
}
