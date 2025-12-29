package moves

import(
	"strings"
)

func SimpleMove(cube map[string][][]string, mv string) map[string][][]string {
	/* todo: implement moves cases */
	switch mv {
	case "U":
		cube = moveU(cube)
	case "U'":
		cube = moveUPrime(cube)
	case "U2":
		cube = moveU2(cube)
	case "R":
		cube = moveR(cube)
	case "R'":
		cube = moveRPrime(cube)
	case "R2":
		cube = moveR2(cube)
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
