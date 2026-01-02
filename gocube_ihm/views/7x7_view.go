package views

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	"simul/gocube/gocube"
	"simul/gocube/gocube_ihm/ihm_components"
)

func Cube7x7View() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Cube 7x7"))
		err := run7x7(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run7x7(window *app.Window) error {
	// theme := material.NewTheme()
	var cube = gocube.GetCube(7)
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)
			ihm_components.SetBackground(gtx)

			// display cube
			ihm_components.Draw7x7Cube(gtx, cube)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
