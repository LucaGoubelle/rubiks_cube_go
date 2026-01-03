package layouts

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func GenMoveButton(gtx layout.Context, theme *material.Theme, moveButton *widget.Clickable) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			btnMove := material.Button(theme, moveButton, "Move")
			return btnMove.Layout(gtx)
		},
	)
}

func GenScrambleButton(gtx layout.Context, theme *material.Theme, scrambleButton *widget.Clickable) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(theme, scrambleButton, "Scramble")
			return btn.Layout(gtx)
		},
	)
}

func GenResetButton(gtx layout.Context, theme *material.Theme, resetButton *widget.Clickable) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			btn2 := material.Button(theme, resetButton, "Reset")
			return btn2.Layout(gtx)
		},
	)
}
