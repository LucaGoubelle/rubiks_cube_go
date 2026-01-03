package views

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"simul/gocube/gocube"
	"simul/gocube/gocube/moves"
	"simul/gocube/gocube_ihm/ihm_components"
	"simul/gocube/gocube_ihm/layouts"
)

func Cube3x3View() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Cube 3x3"))
		err := run3x3(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run3x3(window *app.Window) error {
	theme := material.NewTheme()
	var cube = gocube.GetCube(3)
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
			ihm_components.Draw3x3Cube(gtx, cube)

			if scrambleButton.Clicked(gtx) {
				cube = gocube.ScrambleCube(cube)
			}
			if resetButton.Clicked(gtx) {
				cube = gocube.GetCube(3)
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
				layouts.GenMoveEntry(gtx, theme, &inputMove, &moveTyped),
				layouts.GenSpacer25(),
				layouts.GenMoveButton(gtx, theme, &moveButton),
				layouts.GenSpacer25(),
				layouts.GenScrambleButton(gtx, theme, &scrambleButton),
				layouts.GenSpacer25(),
				layouts.GenResetButton(gtx, theme, &resetButton),
				layouts.GenSpacer25(),
			)
			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
