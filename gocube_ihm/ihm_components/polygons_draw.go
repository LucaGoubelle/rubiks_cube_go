package ihm_components

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/f32"
	"gioui.org/op/paint"
	"gioui.org/op/clip"
)

func drawPolygon(gtx layout.Context, points []f32.Point, col color.NRGBA) {
	if len(points) < 3 {
		return
	}

	var path clip.Path
	path.Begin(gtx.Ops)
	path.MoveTo(points[0])
	for _, p := range points[1:] {
		path.LineTo(p)
	}
	path.Close()

	paint.FillShape(
		gtx.Ops,
		col,
		clip.Outline{Path: path.End()}.Op(),
	)
}
