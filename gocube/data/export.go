package data

import (
	"encoding/json"
	"os"
)

func handleError(err error) {
	if err != nil { panic(err) }
}

func ExportToJSON(cube map[string][][]string, nameFile string) bool {
	file, err := os.Create(nameFile)
	handleError(err)
	jsonString, err := json.MarshalIndent(cube, "", "    ")
	handleError(err)
	length, err := file.WriteString(string(jsonString))
	handleError(err)
	if length > 0 {
		return true
	}
	return false
}
