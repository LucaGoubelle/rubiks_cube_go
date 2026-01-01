package ihm_components

import (
	"gioui.org/layout"
)

func Draw3x3Cube(gtx layout.Context, cube map[string][][]string){
	draw3x3FrontFace(gtx, cube)
	draw3x3UpFace(gtx, cube)
	draw3x3RightFace(gtx, cube)
}
