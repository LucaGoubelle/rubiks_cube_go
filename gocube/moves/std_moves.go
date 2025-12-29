package moves


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

