package gocube_ihm

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/paint"

	"simul/gocube/gocube"
)

func MainView() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Test01"))
		err := runMain(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runMain(window *app.Window) error {
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
				R: 32,
				G: 32,
				B: 32,
				A: 255,
			})

			// display cube
			drawFrontFace(gtx, cube)
			drawUpFace(gtx, cube)
			//todo: implement right faces
			drawRightFace(gtx, cube)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
