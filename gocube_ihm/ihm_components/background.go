package ihm_components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
)

func SetBackground(gtx layout.Context) {
	paint.Fill(gtx.Ops, color.NRGBA{R: 32, G: 32, B: 32, A: 255})
}
