package views

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"simul/gocube/gocube"
	"simul/gocube/gocube_ihm/ihm_components"
	"gioui.org/widget"
	"gioui.org/layout"
	"gioui.org/unit"
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
			//display button + space bottom
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// We insert two rigid elements:
				// First one to hold a button ...
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &scrambleButton, "Scramble")
						return btn.Layout(gtx)
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
