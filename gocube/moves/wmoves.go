package moves

func moveUw(cube map[string][][]string, nbLayers int) map[string][][]string {
	cube["up"] = rotate(cube["up"])
	var size = len(cube["up"])

	var newFront = genEmptyFace(size)
	var newLeft = genEmptyFace(size)
	var newRight = genEmptyFace(size)
	var newBack = genEmptyFace(size)

	for j:=0; j<nbLayers; j++ {
		for i:=0; i<size; i++ {
			newFront[j][i] = cube["right"][j][i]
			newLeft[j][i] = cube["front"][j][i]
			newRight[j][i] = cube["back"][j][i]
			newBack[j][i] = cube["left"][j][i]
		}
	}

	cube["front"] = transfert(cube["front"], newFront)
	cube["left"] = transfert(cube["left"], newLeft)
	cube["right"] = transfert(cube["right"], newRight)
	cube["back"] = transfert(cube["back"], newBack)

	return cube
}

func moveUwPrime(cube map[string][][]string, nbLayers int) map[string][][]string {
	for i:=0;i<3;i++ {
		cube = moveUw(cube, nbLayers)
	}
	return cube
}

func moveUw2(cube map[string][][]string, nbLayers int) map[string][][]string {
	for i:=0;i<2;i++ {
		cube = moveUw(cube, nbLayers)
	}
	return cube
}

func moveDw(cube map[string][][]string, nbLayers int) map[string][][]string {
	cube["down"] = rotate(cube["down"])
	var size = len(cube["down"])

	var newFront = genEmptyFace(size)
	var newLeft = genEmptyFace(size)
	var newRight = genEmptyFace(size)
	var newBack = genEmptyFace(size)

	for j:=0; j<nbLayers; j++ {
		for i:=0; i<size; i++ {
			newFront[size-1-j][i] = cube["left"][size-1-j][i]
			newLeft[size-1-j][i] = cube["back"][size-1-j][i]
			newRight[size-1-j][i] = cube["front"][size-1-j][i]
			newBack[size-1-j][i] = cube["right"][size-1-j][i]
		}
	}
	

	cube["front"] = transfert(cube["front"], newFront)
	cube["left"] = transfert(cube["left"], newLeft)
	cube["right"] = transfert(cube["right"], newRight)
	cube["back"] = transfert(cube["back"], newBack)

	return cube
}

func moveDwPrime(cube map[string][][]string, nbLayers int) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveDw(cube, nbLayers)
	}
	return cube
}

func moveDw2(cube map[string][][]string, nbLayers int) map[string][][]string {
	for i:=0; i<2; i++ {
		cube = moveDw(cube, nbLayers)
	}
	return cube
}

