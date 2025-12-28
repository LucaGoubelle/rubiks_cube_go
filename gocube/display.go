package gocube

import (
	"fmt"
)

var Reset = "\033[0m"
var Red = "\033[41m"
var Green = "\033[42m"
var Yellow = "\033[43m"
var Blue = "\033[44m"
var Magenta = "\033[45m"
var White = "\033[47m"

var colorMap = map[string]string{
	"white":  White + " " + Reset,
	"yellow": Yellow + " " + Reset,
	"blue":   Blue + " " + Reset,
	"red":    Red + " " + Reset,
	"orange": Magenta + " " + Reset,
	"green":  Green + " " + Reset,
}

func printRowUpDown(row []string) string {
	var content = ""
	for i := 0; i < len(row); i++ {
		content += " "
	}
	for _, val := range row {
		content += colorMap[val]
	}
	content += "\n"
	return content
}

func printRowLFRB(rowL []string, rowF []string, rowR []string, rowB []string) string {
	var content = ""
	for _, val := range rowL {
		content += colorMap[val]
	}
	for _, val := range rowF {
		content += colorMap[val]
	}
	for _, val := range rowR {
		content += colorMap[val]
	}
	for _, val := range rowB {
		content += colorMap[val]
	}
	content += "\n"
	return content
}

func PrintCube(cube map[string][][]string) {
	var content = ""
	var size = len(cube["front"])
	for _, val := range cube["up"] {
		content += printRowUpDown(val)
	}
	for i := 0; i < size; i++ {
		content += printRowLFRB(cube["left"][i], cube["front"][i], cube["right"][i], cube["back"][i])
	}
	for _, val := range cube["down"] {
		content += printRowUpDown(val)
	}
	content += "\n\n"
	fmt.Printf(content)
}
