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

func drawUpFace(gtx layout.Context, cube map[string][][]string) {
	drawUpSticker(gtx, f32.Pt(212, 83), cube["up"][2][0], 25)
	drawUpSticker(gtx, f32.Pt(242, 83), cube["up"][2][1], 25)
	drawUpSticker(gtx, f32.Pt(272, 83), cube["up"][2][2], 25)

	drawUpSticker(gtx, f32.Pt(227, 66), cube["up"][1][0], 25)
	drawUpSticker(gtx, f32.Pt(257, 66), cube["up"][1][1], 25)
	drawUpSticker(gtx, f32.Pt(287, 66), cube["up"][1][2], 25)

	drawUpSticker(gtx, f32.Pt(242, 49), cube["up"][0][0], 25)
	drawUpSticker(gtx, f32.Pt(272, 49), cube["up"][0][1], 25)
	drawUpSticker(gtx, f32.Pt(302, 49), cube["up"][0][2], 25)
}

func drawRightFace(gtx layout.Context, cube map[string][][]string) {
	drawRightSticker(gtx, f32.Pt(290,100), cube["right"][0][0], 25)

	drawRightSticker(gtx, f32.Pt(290,130), cube["right"][1][0], 25)

	drawRightSticker(gtx, f32.Pt(290,160), cube["right"][2][0], 25)
}
