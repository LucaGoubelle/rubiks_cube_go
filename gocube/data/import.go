package data

import (
	"encoding/json"
	"os"
)

func ImportFromJSON(path string) map[string][][]string {
	data, err := os.ReadFile(path)
	handleError(err)
	var cube map[string][][]string
	err = json.Unmarshal(data, &cube)
	handleError(err)
	return cube
}
