package layouts

import (
	"image/color"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func GenMoveEntry(gtx layout.Context, theme *material.Theme, inputMove *widget.Editor, moveTyped *string) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			e := material.Editor(theme, inputMove, "")
			e.Color = color.NRGBA{R:255, G:255, B:255, A:255}
			e.HintColor = color.NRGBA{R:255, G:255, B:255, A:255}
			e.Hint = "Enter move here..."
			e.SelectionColor = color.NRGBA{R:64, G:64, B:64, A:255}
			*moveTyped = e.Editor.Text()
			return e.Layout(gtx)
		},
	)
}
