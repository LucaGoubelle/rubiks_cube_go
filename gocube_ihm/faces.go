package gocube_ihm

import (
	"gioui.org/layout"
	"gioui.org/f32"
)

func drawFrontFace(gtx layout.Context, cube map[string][][]string ) {
	drawFrontSticker(gtx, f32.Pt(200,100), cube["front"][0][0], 25)
	drawFrontSticker(gtx, f32.Pt(230,100), cube["front"][0][1], 25)
	drawFrontSticker(gtx, f32.Pt(260,100), cube["front"][0][2], 25)

	drawFrontSticker(gtx, f32.Pt(200,130), cube["front"][1][0], 25)
	drawFrontSticker(gtx, f32.Pt(230,130), cube["front"][1][1], 25)
	drawFrontSticker(gtx, f32.Pt(260,130), cube["front"][1][2], 25)

	drawFrontSticker(gtx, f32.Pt(200,160), cube["front"][2][0], 25)
	drawFrontSticker(gtx, f32.Pt(230,160), cube["front"][2][1], 25)
	drawFrontSticker(gtx, f32.Pt(260,160), cube["front"][2][2], 25)
}
