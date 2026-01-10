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
	case "Uw":
		var size = len(cube["up"])
		if size < 3 {
			cube = moveU(cube)
		} else {
			cube = moveUw(cube, 2)
		}
	case "Uw'":
		var size = len(cube["up"])
		if size < 3 {
			cube = moveUPrime(cube)
		} else {
			cube = moveUwPrime(cube, 2)
		}
	case "Uw2":
		var size = len(cube["up"])
		if size < 3 {
			cube = moveU2(cube)
		} else {
			cube = moveUw2(cube, 2)
		}
	case "R":
		cube = moveR(cube)
	case "R'":
		cube = moveRPrime(cube)
	case "R2":
		cube = moveR2(cube)
	case "L":
		cube = moveL(cube)
	case "L'":
		cube = moveLPrime(cube)
	case "L2":
		cube = moveL2(cube)
	case "Lw":
		var size = len(cube["left"])
		if size < 3 {
			cube = moveL(cube)
		} else {
			cube = moveLw(cube, 2)
		}
	case "Lw'":
		var size = len(cube["left"])
		if size < 3 {
			cube = moveLPrime(cube)
		} else {
			cube = moveLwPrime(cube, 2)
		}
	case "Lw2":
		var size = len(cube["left"])
		if size < 3 {
			cube = moveL2(cube)
		} else {
			cube = moveLw2(cube, 2)
		}
	case "D":
		cube = moveD(cube)
	case "D'":
		cube = moveDPrime(cube)
	case "D2":
		cube = moveD2(cube)
	case "Dw":
		var size = len(cube["down"])
		if size < 3 {
			cube = moveD(cube)
		} else {
			cube = moveDw(cube, 2)
		}
	case "Dw'":
		var size = len(cube["down"])
		if size < 3 {
			cube = moveDPrime(cube)
		} else {
			cube = moveDwPrime(cube, 2)
		}
	case "Dw2":
		var size = len(cube["down"])
		if size < 3 {
			cube = moveD2(cube)
		} else {
			cube = moveDw2(cube, 2)
		}
	case "F":
		cube = moveF(cube)
	case "F'":
		cube = moveFPrime(cube)
	case "F2":
		cube = moveF2(cube)
	case "x":
		cube = cube
	case "y":
		cube = moveY(cube)
	case "y'":
		cube = moveYPrime(cube)
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
