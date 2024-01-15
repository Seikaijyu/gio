// SPDX-License-Identifier: Unlicense OR MIT

package widget

import (
	"image"
	"image/color"

	"github.com/Seikaijyu/gio/layout"
	"github.com/Seikaijyu/gio/op/clip"
	"github.com/Seikaijyu/gio/op/paint"
	"github.com/Seikaijyu/gio/unit"
)

// Border lays out a widget and draws a border inside it.
type Border struct {
	Color        color.NRGBA
	CornerRadius unit.Dp
	Width        unit.Dp
}

func (b Border) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	dims := w(gtx)
	sz := dims.Size

	rr := gtx.Dp(b.CornerRadius)
	width := gtx.Dp(b.Width)
	whalf := (width + 1) / 2
	sz.X -= whalf * 2
	sz.Y -= whalf * 2

	r := image.Rectangle{Max: sz}
	r = r.Add(image.Point{X: whalf, Y: whalf})

	paint.FillShape(gtx.Ops,
		b.Color,
		clip.Stroke{
			Path:  clip.UniformRRect(r, rr).Path(gtx.Ops),
			Width: float32(width),
		}.Op(),
	)

	return dims
}
