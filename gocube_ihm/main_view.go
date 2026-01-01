package gocube_ihm

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
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
			//todo: implement up / right faces
			drawFrontSticker(gtx, f32.Pt(200,100), cube["front"][0][0], 25)
			drawFrontSticker(gtx, f32.Pt(230,100), cube["front"][0][1], 25)
			drawFrontSticker(gtx, f32.Pt(260,100), cube["front"][0][2], 25)

			drawFrontSticker(gtx, f32.Pt(200,130), cube["front"][1][0], 25)
			drawFrontSticker(gtx, f32.Pt(230,130), cube["front"][1][1], 25)
			drawFrontSticker(gtx, f32.Pt(260,130), cube["front"][1][2], 25)

			drawFrontSticker(gtx, f32.Pt(200,160), cube["front"][2][0], 25)
			drawFrontSticker(gtx, f32.Pt(230,160), cube["front"][2][1], 25)
			drawFrontSticker(gtx, f32.Pt(260,160), cube["front"][2][2], 25)


			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

