package ihm_components

import (
	"gioui.org/f32"
	"gioui.org/layout"
)

func _drawFront(gtx layout.Context, face [][]string, coord f32.Point, size int, offset int) {
	for i := 0; i < len(face); i++ {
		for j := 0; j < len(face); j++ {
			xActual := coord.X + float32(size)*float32(j) + float32(offset)*float32(j)
			yActual := coord.Y + float32(size)*float32(i) + float32(offset)*float32(i)

			drawFrontSticker(gtx, f32.Pt(xActual, yActual), face[i][j], size)
		}
	}
}

func _drawRight(gtx layout.Context, face [][]string, coord f32.Point, size int, half int, offset int) {
	faceLen := len(face)
	x := coord.X
	y := coord.Y

	for i := 0; i < faceLen; i++ {
		var totalOffsetI = offset * i
		var totalHalf = half * i
		var allSize = size * faceLen
		var allOffset = offset * (faceLen + i)
		var xFinal = x + float32(allSize) + float32(totalHalf) + float32(allOffset)

		for j := 0; j < faceLen; j++ {
			var totalOffsetJ = offset * j
			var totalSize = size * j
			var y0 = y + float32(totalSize) + float32(totalOffsetJ)
			var yFinal = float32(y0) - float32(totalHalf) - float32(totalOffsetI)

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

func draw4x4UpFace(gtx layout.Context, face [][]string, coord f32.Point, size int, half int, offset int) {
	var faceLen = len(face)

	for i := 0; i < faceLen; i++ {
		var x0 = coord.X + float32(size)*(float32(i)+float32(0)) + float32(offset)*(float32(i)+float32(0)) + float32(half)
		var x1 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(1))
		var x2 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(2)) + float32(half)
		var x3 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(3))

		var y0 = coord.Y - float32(half)*float32(1) - float32(offset)*float32(1)
		var y1 = coord.Y - float32(half)*float32(2) - float32(offset)*float32(2)
		var y2 = coord.Y - float32(half)*float32(3) - float32(offset)*float32(3)
		var y3 = coord.Y - float32(half)*float32(4) - float32(offset)*float32(4)

		drawUpSticker(gtx, f32.Pt(x0, y0), face[3][i], size)
		drawUpSticker(gtx, f32.Pt(x1, y1), face[2][i], size)
		drawUpSticker(gtx, f32.Pt(x2, y2), face[1][i], size)
		drawUpSticker(gtx, f32.Pt(x3, y3), face[0][i], size)
	}
}

func draw5x5UpFace(gtx layout.Context, face [][]string, coord f32.Point, size int, half int, offset int) {
	var faceLen = len(face)

	for i := 0; i < faceLen; i++ {
		var x0 = coord.X + float32(size)*(float32(i)+float32(0)) + float32(offset)*(float32(i)+float32(0)) + float32(half)
		var x1 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(1))
		var x2 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(2)) + float32(half)
		var x3 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(3))
		var x4 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(4)) + float32(half)

		var y0 = coord.Y - float32(half)*float32(1) - float32(offset)*float32(1)
		var y1 = coord.Y - float32(half)*float32(2) - float32(offset)*float32(2)
		var y2 = coord.Y - float32(half)*float32(3) - float32(offset)*float32(3)
		var y3 = coord.Y - float32(half)*float32(4) - float32(offset)*float32(4)
		var y4 = coord.Y - float32(half)*float32(5) - float32(offset)*float32(5)

		drawUpSticker(gtx, f32.Pt(x0, y0), face[4][i], size)
		drawUpSticker(gtx, f32.Pt(x1, y1), face[3][i], size)
		drawUpSticker(gtx, f32.Pt(x2, y2), face[2][i], size)
		drawUpSticker(gtx, f32.Pt(x3, y3), face[1][i], size)
		drawUpSticker(gtx, f32.Pt(x4, y4), face[0][i], size)
	}
}

func draw6x6UpFace(gtx layout.Context, face [][]string, coord f32.Point, size int, half int, offset int) {
	var faceLen = len(face)

	for i := 0; i < faceLen; i++ {
		var x0 = coord.X + float32(size)*(float32(i)+float32(0)) + float32(offset)*(float32(i)+float32(0)) + float32(half)
		var x1 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(1))
		var x2 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(2)) + float32(half)
		var x3 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(3))
		var x4 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(4)) + float32(half)
		var x5 = coord.X + float32(size)*(float32(i)+float32(3)) + float32(offset)*(float32(i)+float32(5))

		var y0 = coord.Y - float32(half)*float32(1) - float32(offset)*float32(1)
		var y1 = coord.Y - float32(half)*float32(2) - float32(offset)*float32(2)
		var y2 = coord.Y - float32(half)*float32(3) - float32(offset)*float32(3)
		var y3 = coord.Y - float32(half)*float32(4) - float32(offset)*float32(4)
		var y4 = coord.Y - float32(half)*float32(5) - float32(offset)*float32(5)
		var y5 = coord.Y - float32(half)*float32(6) - float32(offset)*float32(6)

		drawUpSticker(gtx, f32.Pt(x0, y0), face[5][i], size)
		drawUpSticker(gtx, f32.Pt(x1, y1), face[4][i], size)
		drawUpSticker(gtx, f32.Pt(x2, y2), face[3][i], size)
		drawUpSticker(gtx, f32.Pt(x3, y3), face[2][i], size)
		drawUpSticker(gtx, f32.Pt(x4, y4), face[1][i], size)
		drawUpSticker(gtx, f32.Pt(x5, y5), face[0][i], size)
	}
}

func draw7x7UpFace(gtx layout.Context, face [][]string, coord f32.Point, size int, half int, offset int) {
	var faceLen = len(face)

	for i := 0; i < faceLen; i++ {
		var x0 = coord.X + float32(size)*(float32(i)+float32(0)) + float32(offset)*(float32(i)+float32(0)) + float32(half)
		var x1 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(1))
		var x2 = coord.X + float32(size)*(float32(i)+float32(1)) + float32(offset)*(float32(i)+float32(2)) + float32(half)
		var x3 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(3))
		var x4 = coord.X + float32(size)*(float32(i)+float32(2)) + float32(offset)*(float32(i)+float32(4)) + float32(half)
		var x5 = coord.X + float32(size)*(float32(i)+float32(3)) + float32(offset)*(float32(i)+float32(5))
		var x6 = coord.X + float32(size)*(float32(i)+float32(3)) + float32(offset)*(float32(i)+float32(6)) + float32(half)

		var y0 = coord.Y - float32(half)*float32(1) - float32(offset)*float32(1)
		var y1 = coord.Y - float32(half)*float32(2) - float32(offset)*float32(2)
		var y2 = coord.Y - float32(half)*float32(3) - float32(offset)*float32(3)
		var y3 = coord.Y - float32(half)*float32(4) - float32(offset)*float32(4)
		var y4 = coord.Y - float32(half)*float32(5) - float32(offset)*float32(5)
		var y5 = coord.Y - float32(half)*float32(6) - float32(offset)*float32(6)
		var y6 = coord.Y - float32(half)*float32(7) - float32(offset)*float32(7)

		drawUpSticker(gtx, f32.Pt(x0, y0), face[6][i], size)
		drawUpSticker(gtx, f32.Pt(x1, y1), face[5][i], size)
		drawUpSticker(gtx, f32.Pt(x2, y2), face[4][i], size)
		drawUpSticker(gtx, f32.Pt(x3, y3), face[3][i], size)
		drawUpSticker(gtx, f32.Pt(x4, y4), face[2][i], size)
		drawUpSticker(gtx, f32.Pt(x5, y5), face[1][i], size)
		drawUpSticker(gtx, f32.Pt(x6, y6), face[0][i], size)
	}
}
