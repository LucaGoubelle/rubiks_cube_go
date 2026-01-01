package ihm_components

import (
	"gioui.org/layout"
	"gioui.org/f32"
)

func Draw3x3Cube(gtx layout.Context, cube map[string][][]string){
	_drawFront(gtx, cube["front"], f32.Pt(200,100), 25, 5)
	draw3x3UpFace(gtx, cube)
	_drawRight(gtx, cube["right"], f32.Pt(200,100), 25, 12, 5)
}
