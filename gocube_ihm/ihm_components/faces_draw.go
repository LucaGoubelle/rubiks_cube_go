package ihm_components

import (
	"gioui.org/layout"
	"gioui.org/f32"
)

func _drawFront(gtx layout.Context, face [][]string, coord f32.Point, size int, offset int) {
	for i:=0; i<len(face); i++ {
		for j:=0; j<len(face); j++ {
			xActual := coord.X + float32(size) * float32(j) + float32(offset) * float32(j)
			yActual := coord.Y + float32(size) * float32(i) + float32(offset) * float32(i)

			drawFrontSticker(gtx, f32.Pt(xActual, yActual), face[i][j], size)
		}
	}
}

func _drawRight(gtx layout.Context, face [][]string, coord f32.Point, size int, half int, offset int) {
	faceLen := len(face)
	x := coord.X
	y := coord.Y

	for i:=0; i<faceLen; i++ {
		var totalOffsetI = offset * i
		var totalHalf = half * i
		var allSize = size * faceLen
		var allOffset = offset * (faceLen+i)
		var xFinal = x + float32(allSize) + float32(totalHalf) + float32(allOffset)

		for j:=0; j<faceLen; j++ {
			var totalOffsetJ = offset * j
			var totalSize = size * j
			var y0 = y + float32(totalSize) + float32(totalOffsetJ)
			var yFinal = float32(y0)-float32(totalHalf)-float32(totalOffsetI)

			drawRightSticker(gtx, f32.Pt(xFinal, yFinal), face[j][i], size)
		}
	}
}

func draw3x3UpFace(gtx layout.Context, cube map[string][][]string) {
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
