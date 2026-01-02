package ihm_components

import (
	"gioui.org/f32"
	"gioui.org/layout"
)

func Draw3x3Cube(gtx layout.Context, cube map[string][][]string) {
	_drawFront(gtx, cube["front"], f32.Pt(200, 100), 25, 5)
	draw3x3UpFace(gtx, cube)
	_drawRight(gtx, cube["right"], f32.Pt(200, 100), 25, 12, 5)
}

func Draw4x4Cube(gtx layout.Context, cube map[string][][]string) {
	_drawFront(gtx, cube["front"], f32.Pt(200, 100), 25, 5)
	draw4x4UpFace(gtx, cube["up"], f32.Pt(200, 100), 25, 12, 5)
	_drawRight(gtx, cube["right"], f32.Pt(200, 100), 25, 12, 5)
}

func Draw5x5Cube(gtx layout.Context, cube map[string][][]string) {
	_drawFront(gtx, cube["front"], f32.Pt(200, 100), 25, 5)
	_drawRight(gtx, cube["right"], f32.Pt(200, 100), 25, 12, 5)
}

func Draw6x6Cube(gtx layout.Context, cube map[string][][]string) {
	_drawFront(gtx, cube["front"], f32.Pt(200, 100), 20, 5)
	_drawRight(gtx, cube["right"], f32.Pt(200, 100), 20, 10, 5)
}

func Draw7x7Cube(gtx layout.Context, cube map[string][][]string) {
	_drawFront(gtx, cube["front"], f32.Pt(200, 100), 15, 5)
	_drawRight(gtx, cube["right"], f32.Pt(200, 100), 15, 7, 5)
}
