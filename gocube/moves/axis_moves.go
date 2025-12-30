package moves

func moveY(cube map[string][][]string) map[string][][]string {
	cube["up"] = rotate(cube["up"])
	cube["down"] = rotateAsync(cube["down"])

	var newFront = copyFace(cube["right"])
	var newRight = copyFace(cube["back"])
	var newLeft = copyFace(cube["front"])
	var newBack = copyFace(cube["left"])

	cube["front"] = transfert(cube["front"], newFront)
	cube["right"] = transfert(cube["right"], newRight)
	cube["left"] = transfert(cube["left"], newLeft)
	cube["back"] = transfert(cube["back"], newBack)
	
	return cube
}

func moveYPrime(cube map[string][][]string) map[string][][]string {
	for i:=0; i<3; i++ {
		cube = moveY(cube)
	}
	return cube
}
