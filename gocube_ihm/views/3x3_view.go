package views

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"

	"simul/gocube/gocube"
	"simul/gocube/gocube_ihm/ihm_components"
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
	// theme := material.NewTheme()
	var cube = gocube.GetCube(3)
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			paint.Fill(gtx.Ops, color.NRGBA{
				R: 32, G: 32, B: 32, A: 255,
			})

			// display cube
			ihm_components.Draw3x3Cube(gtx, cube)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
