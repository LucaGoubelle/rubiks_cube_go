package layouts

import (
	"gioui.org/unit"
	"gioui.org/layout"
)

func GenSpacer25() layout.FlexChild {
	return layout.Rigid(
		// The height of the spacer is 25 Device independent pixels
		layout.Spacer{Height: unit.Dp(25)}.Layout,
	)
}
