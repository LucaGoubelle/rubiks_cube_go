package gocube

import(
	"strings"
)

func SimpleMove(cube map[string][][]string, mv string) map[string][][]string {
	/* todo: implement moves cases */
	switch mv {
	case "U":
		cube = cube
	case "R":
		cube = cube
	case "L":
		cube = cube
	case "D":
		cube = cube
	case "F":
		cube = cube
	case "x":
		cube = cube
	case "y":
		cube = cube
	case "z":
		cube = cube
	}
	return cube
}

func MultiMoves(cube map[string][][]string, mvs string) map[string][][]string {
	var mvsList = strings.Split(mvs, " ")
	for _,val := range mvsList {
		cube = SimpleMove(cube, val)
	}
	return cube
}
