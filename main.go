package main

import (
	"fmt"

	"simul/gocube/gocube"
)

func main() {
	var cube = gocube.GetCube(3)
	gocube.PrintCube(cube)
	fmt.Println(cube["front"][0][0])
}
