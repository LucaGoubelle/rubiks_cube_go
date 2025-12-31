package gocube

import (
    "math/rand/v2"
	"simul/gocube/gocube/moves"
)

func ScrambleCube(cube map[string][][]string) map[string][][]string {
	scrambles := []string{
		"F' L2 R2 D2 L2 D R' F' U B L R2 U' L U L' R D' L F2 B' D' L' D U2", 
		"B2 D2 U' F L2 R' F L R F2 D2 L B' U2 D' R U' B U2 R U2 D' L F' B", 
		"B L D L2 B2 L2 U' F2 D' U' L' B D' F2 D U R' F L F2 R' L' D F U", 
		"R2 L' U2 R2 B2 R' B' D R2 L2 U D' B' D F2 L' F' D' F' U' F U2 R2 U R2",
		"L' B F U' F L B2 U2 D R' D L2 U F' D U2 L' D B2 U R' L D2 F R2",
	}
	var randint = rand.IntN(len(scrambles)-1)
	var scrambling = scrambles[randint]
	cube = moves.MultiMoves(cube, scrambling)
	return cube
}
