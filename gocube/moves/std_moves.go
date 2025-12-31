package moves

/* #######################
##### U / D MOVES ########
##########################*/

func moveU(cube map[string][][]string) map[string][][]string {
	cube["up"] = rotate(cube["up"])
	var size = len(cube["up"])

	var newFront = genEmptyFace(size)
	var newLeft = genEmptyFace(size)
	var newRight = genEmptyFace(size)
	var newBack = genEmptyFace(size)

	for i:=0; i<size; i++ {
		newFront[0][i] = cube["right"][0][i]
		newLeft[0][i] = cube["front"][0][i]
		newRight[0][i] = cube["back"][0][i]
		newBack[0][i] = cube["left"][0][i]
	}

	cube["front"] = transfert(cube["front"], newFront)
	cube["left"] = transfert(cube["left"], newLeft)
	cube["right"] = transfert(cube["right"], newRight)
	cube["back"] = transfert(cube["back"], newBack)

	return cube
}

func moveUPrime(cube map[string][][]string) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveU(cube)
	}
	return cube
}

func moveU2(cube map[string][][]string) map[string][][]string {
	for i:=0; i<2; i++ {
		cube = moveU(cube)
	}
	return cube
}

func moveD(cube map[string][][]string) map[string][][]string {
	cube["down"] = rotate(cube["down"])
	var size = len(cube["down"])

	var newFront = genEmptyFace(size)
	var newLeft = genEmptyFace(size)
	var newRight = genEmptyFace(size)
	var newBack = genEmptyFace(size)

	for i:=0; i<size; i++ {
		newFront[size-1][i] = cube["left"][size-1][i]
		newLeft[size-1][i] = cube["back"][size-1][i]
		newRight[size-1][i] = cube["front"][size-1][i]
		newBack[size-1][i] = cube["right"][size-1][i]
	}

	cube["front"] = transfert(cube["front"], newFront)
	cube["left"] = transfert(cube["left"], newLeft)
	cube["right"] = transfert(cube["right"], newRight)
	cube["back"] = transfert(cube["back"], newBack)

	return cube
}

func moveDPrime(cube map[string][][]string) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveD(cube)
	}
	return cube
}

func moveD2(cube map[string][][]string) map[string][][]string {
	for i:=0; i<2; i++ {
		cube = moveD(cube)
	}
	return cube
}

/* ############################
##### L / R MOVES  ############
###############################*/

func moveL(cube map[string][][]string) map[string][][]string {
	cube["left"] = rotate(cube["left"])
	var size = len(cube["left"])

	var newUp = genEmptyFace(size)
	var newFront = genEmptyFace(size)
	var newDown = genEmptyFace(size)
	var newBack = genEmptyFace(size)

	for i:=0; i<size; i++ {
		newFront[i][0] = cube["up"][i][0]
		newDown[i][0] = cube["front"][i][0]
		newBack[i][0] = cube["down"][i][0]
		newUp[i][size-1] = cube["back"][i][size-1]
	}

	cube["front"] = transfert(cube["front"], newFront)
	cube["up"] = transfert(cube["up"], rotateTwice(newUp))
	cube["down"] = transfert(cube["down"], newDown)
	cube["back"] = transfert(cube["back"], rotateTwice(newBack))

	return cube
}

func moveLPrime(cube map[string][][]string) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveL(cube)
	}
	return cube
}

func moveL2(cube map[string][][]string) map[string][][]string {
	for i:=0; i<2; i++ {
		cube = moveL(cube)
	}
	return cube
}

func moveR(cube map[string][][]string) map[string][][]string {
	cube["right"] = rotate(cube["right"])
	var size = len(cube["right"])

	var newFront = genEmptyFace(size)
	var newUp = genEmptyFace(size)
	var newBack = genEmptyFace(size)
	var newDown = genEmptyFace(size)

	for i:=0; i<size; i++ {
		newFront[i][size-1] = cube["down"][i][size-1]
		newUp[i][size-1] = cube["front"][i][size-1]
		newBack[i][size-1] = cube["up"][i][size-1]
		newDown[i][0] = cube["back"][i][0]
	}

	cube["front"] = transfert(cube["front"], newFront)
	cube["up"] = transfert(cube["up"], newUp)
	cube["back"] = transfert(cube["back"], rotateTwice(newBack))
	cube["down"] = transfert(cube["down"], rotateTwice(newDown))

	return cube
}

func moveRPrime(cube map[string][][]string) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveR(cube)
	}
	return cube
}

func moveR2(cube map[string][][]string) map[string][][]string {
	for i:=0; i<2; i++ {
		cube = moveR(cube)
	}
	return cube
}

func moveF(cube map[string][][]string) map[string][][]string {
	cube["front"] = rotate(cube["front"])
	var size = len(cube["front"])

	var newUp = genEmptyFace(size)
	var newLeft = genEmptyFace(size)
	var newRight = genEmptyFace(size)
	var newDown = genEmptyFace(size)

	for i:=0; i<size; i++ {
		newUp[i][size-1] = cube["left"][i][size-1]
		newLeft[0][i] = cube["down"][0][i]
		newRight[size-1][i] = cube["up"][size-1][i]
		newDown[i][0] = cube["right"][i][0]
	}

	cube["up"] = transfert(cube["up"], rotate(newUp))
	cube["left"] = transfert(cube["left"], rotate(newLeft))
	cube["right"] = transfert(cube["right"], rotate(newRight))
	cube["down"] = transfert(cube["down"], rotate(newDown))

	return cube
}

func moveFPrime(cube map[string][][]string) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveF(cube)
	}
	return cube
}

func moveF2(cube map[string][][]string) map[string][][]string {
	for i:=0; i<2; i++ {
		cube = moveF(cube)
	}
	return cube
}
