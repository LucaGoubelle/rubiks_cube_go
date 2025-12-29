package moves


func moveU(cube map[string][][]string) map[string][][]string {
	/* todo: implement this */
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
