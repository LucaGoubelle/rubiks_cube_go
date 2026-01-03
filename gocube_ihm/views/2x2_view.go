package views

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"simul/gocube/gocube"
	"simul/gocube/gocube/moves"
	"simul/gocube/gocube_ihm/ihm_components"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Cube2x2View() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Cube 2x2"))
		err := run2x2(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run2x2(window *app.Window) error {
	theme := material.NewTheme()
	var cube = gocube.GetCube(2)
	var ops op.Ops
	var scrambleButton widget.Clickable
	var resetButton widget.Clickable
	var moveButton widget.Clickable
	var inputMove widget.Editor
	var moveTyped string
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)
			ihm_components.SetBackground(gtx)

			// display cube
			ihm_components.Draw2x2Cube(gtx, cube)

			if scrambleButton.Clicked(gtx) {
				cube = gocube.ScrambleCube(cube)
			}
			if resetButton.Clicked(gtx) {
				cube = gocube.GetCube(2)
			}
			if moveButton.Clicked(gtx) {
				cube = moves.MultiMoves(cube, moveTyped)
			}
			//display button + space bottom
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						e := material.Editor(theme, &inputMove, "")
						e.Color = color.NRGBA{R:255, G:255, B:255, A:255}
						e.HintColor = color.NRGBA{R:255, G:255, B:255, A:255}
						e.Hint = "Enter move here..."
						e.SelectionColor = color.NRGBA{R:64, G:64, B:64, A:255}
						moveTyped = e.Editor.Text()
						return e.Layout(gtx)
					},
				),
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btnMove := material.Button(theme, &moveButton, "Move")
						return btnMove.Layout(gtx)
					},
				),
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
				// We insert two rigid elements:
				// First one to hold a button ...
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &scrambleButton, "Scramble")
						return btn.Layout(gtx)
					},
				),
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn2 := material.Button(theme, &resetButton, "Reset")
						return btn2.Layout(gtx)
					},
				),
				// ... then one to hold an empty spacer
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)
			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
