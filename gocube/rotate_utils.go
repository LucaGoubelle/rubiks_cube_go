package gocube

func genEmptyFace(size int) [][]string {
	var face = make([][]string, size)
	for i := 0; i < size; i++ {
		var row = make([]string, size)
		for j := 0; j < size; j++ {
			row[j] = ""
		}
		face[i] = row
	}
	return face
}

func copyFace(face [][]string) [][]string {
	var newFace = make([][]string, len(face))
	for i := 0; i < len(face); i++ {
		var row = make([]string, len(face))
		for j := 0; j < len(face); j++ {
			row[j] = face[i][j]
		}
		newFace[i] = row
	}
	return newFace
}

func rotate(face [][]string ) [][]string {
	var size = len(face)
	var newFace = genEmptyFace(size)
	for i:=0; i<size; i++ {
		var cnt = size-1
		for j:=0; j<size; j++ {
			newFace[i][j] = face[cnt][i]
			cnt--
		}
	}
	return newFace
}

func rotateAsync(face [][]string) [][]string {
	var size = len(face)
	var newFace = genEmptyFace(size)
	var cnt = size-1
	for i:=0; i<size; i++ {
		for j:=0; j<size; j++ {
			newFace[i][j] = face[j][cnt]
		}
		cnt--
	}
	return newFace
}

func rotateTwice(face [][]string) [][]string {
	var newFace = rotate(face)
	return rotate(newFace)
}

func transfert(face [][]string, newFace [][]string) [][]string {
	for i:=0; i<len(face); i++ {
		for j:=0; j<len(face); j++ {
			if newFace[i][j] != "" {
				face[i][j] = newFace[i][j]
			} else {
				face[i][j] = face[i][j]
			}
		}
	}
	return face
}
