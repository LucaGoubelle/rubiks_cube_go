package gominx

func _getKilominxFace(elem string) []string {
	var face = []string{
		elem, elem, elem, elem, elem,
	}
	return face
}

func _getMegaminxFace(elem string) [][]string {
	var face = [][]string{{
		elem, elem, elem, elem, elem,
		elem, elem, elem, elem, elem,
	}, {elem}}
	return face
}

func GetKilominx() map[string][]string {
	var minx = make(map[string][]string)
	minx["up"] = _getKilominxFace("gray")
	minx["front"] = _getKilominxFace("magenta")
	minx["left"] = _getKilominxFace("lime")
	minx["right"] = _getKilominxFace("beige")
	minx["downLeft"] = _getKilominxFace("blue")
	minx["downRight"] = _getKilominxFace("red")

	minx["absLeft"] = _getKilominxFace("yellow")
	minx["absRight"] = _getKilominxFace("green")
	minx["backLeft"] = _getKilominxFace("orange")
	minx["backRight"] = _getKilominxFace("cyan")
	minx["back"] = _getKilominxFace("purple")
	minx["down"] = _getKilominxFace("white")
	return minx
}

func GetMegaminx() map[string][][]string {
	var minx = make(map[string][][]string)
	minx["up"] = _getMegaminxFace("gray")
	minx["front"] = _getMegaminxFace("magenta")
	minx["left"] = _getMegaminxFace("lime")
	minx["right"] = _getMegaminxFace("beige")
	minx["downLeft"] = _getMegaminxFace("blue")
	minx["downRight"] = _getMegaminxFace("red")

	minx["absLeft"] = _getMegaminxFace("yellow")
	minx["absRight"] = _getMegaminxFace("green")
	minx["backLeft"] = _getMegaminxFace("orange")
	minx["backRight"] = _getMegaminxFace("cyan")
	minx["back"] = _getMegaminxFace("purple")
	minx["down"] = _getMegaminxFace("white")
	return minx
}
