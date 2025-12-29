package gocube

func _getFace(size int, elem string) [][]string {
	var face = make([][]string, size)
	for i := 0; i < size; i++ {
		var row = make([]string, size)
		for j := 0; j < size; j++ {
			row[j] = elem
		}
		face[i] = row
	}
	return face
}

func GetCube(size int) map[string][][]string {
	var cube = make(map[string][][]string)
	cube["back"] = _getFace(size, "green")
	cube["up"] = _getFace(size, "yellow")
	cube["front"] = _getFace(size, "blue")
	cube["left"] = _getFace(size, "orange")
	cube["right"] = _getFace(size, "red")
	cube["down"] = _getFace(size, "white")
	return cube
}
