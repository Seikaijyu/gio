// SPDX-License-Identifier: Unlicense OR MIT

package widget

import (
	"image"

	"github.com/Seikaijyu/gio/f32"
	"github.com/Seikaijyu/gio/layout"
	"github.com/Seikaijyu/gio/op"
	"github.com/Seikaijyu/gio/op/clip"
	"github.com/Seikaijyu/gio/op/paint"
	"github.com/Seikaijyu/gio/unit"
)

// Image is a widget that displays an image.
type Image struct {
	// Src is the image to display.
	Src paint.ImageOp
	// Fit specifies how to scale the image to the constraints.
	// By default it does not do any scaling.
	Fit Fit
	// Position specifies where to position the image within
	// the constraints.
	Position layout.Direction
	// Scale is the factor used for converting image pixels to dp.
	// If Scale is zero it defaults to 1.
	//
	// To map one image pixel to one output pixel, set Scale to 1.0 / gtx.Metric.PxPerDp.
	Scale float32
}

func (im Image) Layout(gtx layout.Context) layout.Dimensions {
	scale := im.Scale
	if scale == 0 {
		scale = 1
	}

	size := im.Src.Size()
	wf, hf := float32(size.X), float32(size.Y)
	w, h := gtx.Dp(unit.Dp(wf*scale)), gtx.Dp(unit.Dp(hf*scale))

	dims, trans := im.Fit.scale(gtx.Constraints, im.Position, layout.Dimensions{Size: image.Pt(w, h)})
	defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()

	pixelScale := scale * gtx.Metric.PxPerDp
	trans = trans.Mul(f32.Affine2D{}.Scale(f32.Point{}, f32.Pt(pixelScale, pixelScale)))
	defer op.Affine(trans).Push(gtx.Ops).Pop()

	im.Src.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return dims
}
