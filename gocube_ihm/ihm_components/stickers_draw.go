package ihm_components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/f32"
)

var mapStickerColors = map[string]color.NRGBA {
	"blue": color.NRGBA{B: 255, A:255},
	"green": color.NRGBA{G: 255, A:255},
	"red": color.NRGBA{R: 255, A:255},
	"yellow": color.NRGBA{R: 255, G: 255, A:255},
	"orange": color.NRGBA{R: 255, G: 190, A:255},
	"white": color.NRGBA{R: 255, G: 255, B:255, A:255},
	"black": color.NRGBA{R: 0, G: 0, B: 0, A:255},
}

func drawFrontSticker(gtx layout.Context, coord f32.Point, col string, size int) {
	colors, ok := mapStickerColors[col]
	if ok {} else {
		colors = mapStickerColors["black"]
	}
	points := []f32.Point{
		f32.Pt(coord.X, coord.Y),
		f32.Pt(coord.X+float32(size), coord.Y),
		f32.Pt(coord.X+float32(size), coord.Y+float32(size)),
		f32.Pt(coord.X, coord.Y+float32(size)),
	}
	drawPolygon(gtx, points, colors)
}

func drawUpSticker(gtx layout.Context, coord f32.Point, col string, size int) {
	colors, ok := mapStickerColors[col]
	if ok {} else {
		colors = mapStickerColors["black"]
	}
	half := int(size / 2)
	points := []f32.Point{
		f32.Pt(coord.X, coord.Y),
		f32.Pt(coord.X+float32(size), coord.Y),
		f32.Pt(coord.X+float32(size)-float32(half), coord.Y+float32(half)),
		f32.Pt(coord.X-float32(half), coord.Y+float32(half)),
	}
	drawPolygon(gtx, points, colors)
}

func drawRightSticker(gtx layout.Context, coord f32.Point, col string, size int){
	colors, ok := mapStickerColors[col]
	if ok {} else {
		colors = mapStickerColors["black"]
	}
	half := int(size / 2)
	points := []f32.Point{
		f32.Pt(coord.X, coord.Y),
		f32.Pt(coord.X+float32(half), coord.Y-float32(half)),
		f32.Pt(coord.X+float32(half), coord.Y+float32(half)),
		f32.Pt(coord.X, coord.Y+float32(size)),
	}
	drawPolygon(gtx, points, colors)
}
