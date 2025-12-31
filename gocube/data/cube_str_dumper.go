package data

var mapChar = map[string]string {
	"white":  "W",
	"yellow": "Y",
	"blue":   "B",
	"red":    "R",
	"orange": "O",
	"green":  "G",
}

func dumpFace(face [][]string) string {
	var content string = ""
	for _, row := range face {
		for _,elem := range row {
			val, ok := mapChar[elem]
			if ok {
				content += val
			} else {
				content += "?"
			}
		}
	}
	return content
}

func DumpStr(cube map[string][][]string) string {
	var content string = ""
	content += dumpFace(cube["back"]) + "_"
	content += dumpFace(cube["up"]) + "_"
	content += dumpFace(cube["front"]) + "_"
	content += dumpFace(cube["left"]) + "_"
	content += dumpFace(cube["right"]) + "_"
	content += dumpFace(cube["down"])
	return content
}
