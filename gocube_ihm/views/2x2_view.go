package views

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"simul/gocube/gocube"
	"simul/gocube/gocube/moves"
	"simul/gocube/gocube_ihm/ihm_components"
	"simul/gocube/gocube_ihm/layouts"

	"gioui.org/layout"
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
				layouts.GenMoveEntry(gtx, theme, &inputMove, &moveTyped),
				layouts.GenSpacer25(),
				layouts.GenMoveButton(gtx, theme, &moveButton),
				layouts.GenSpacer25(),
				layouts.GenScrambleButton(gtx, theme, &scrambleButton),
				layouts.GenSpacer25(),
				layouts.GenResetButton(gtx, theme, &resetButton),
				layouts.GenSpacer25(),
			)

			e.Frame(gtx.Ops)
		}
	}
}
