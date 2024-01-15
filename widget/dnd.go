package widget

import (
	"io"

	"github.com/Seikaijyu/gio/f32"
	"github.com/Seikaijyu/gio/gesture"
	"github.com/Seikaijyu/gio/io/pointer"
	"github.com/Seikaijyu/gio/io/transfer"
	"github.com/Seikaijyu/gio/layout"
	"github.com/Seikaijyu/gio/op"
	"github.com/Seikaijyu/gio/op/clip"
)

// Draggable makes a widget draggable.
type Draggable struct {
	// Type contains the MIME type and matches transfer.SourceOp.
	Type string

	handle struct{}
	drag   gesture.Drag
	click  f32.Point
	pos    f32.Point
}

func (d *Draggable) Layout(gtx layout.Context, w, drag layout.Widget) layout.Dimensions {
	if gtx.Queue == nil {
		return w(gtx)
	}
	dims := w(gtx)

	stack := clip.Rect{Max: dims.Size}.Push(gtx.Ops)
	d.drag.Add(gtx.Ops)
	transfer.SourceOp{
		Tag:  &d.handle,
		Type: d.Type,
	}.Add(gtx.Ops)
	stack.Pop()

	if drag != nil && d.drag.Pressed() {
		rec := op.Record(gtx.Ops)
		op.Offset(d.pos.Round()).Add(gtx.Ops)
		drag(gtx)
		op.Defer(gtx.Ops, rec.Stop())
	}

	return dims
}

// Dragging returns whether d is being dragged.
func (d *Draggable) Dragging() bool {
	return d.drag.Dragging()
}

// Update the draggable and returns the MIME type for which the Draggable was
// requested to offer data, if any
func (d *Draggable) Update(gtx layout.Context) (mime string, requested bool) {
	pos := d.pos
	for _, ev := range d.drag.Update(gtx.Metric, gtx.Queue, gesture.Both) {
		switch ev.Kind {
		case pointer.Press:
			d.click = ev.Position
			pos = f32.Point{}
		case pointer.Drag, pointer.Release:
			pos = ev.Position.Sub(d.click)
		}
	}
	d.pos = pos

	for _, ev := range gtx.Queue.Events(&d.handle) {
		if e, ok := ev.(transfer.RequestEvent); ok {
			return e.Type, true
		}
	}
	return "", false
}

// Offer the data ready for a drop. Must be called after being Requested.
// The mime must be one in the requested list.
func (d *Draggable) Offer(ops *op.Ops, mime string, data io.ReadCloser) {
	transfer.OfferOp{
		Tag:  &d.handle,
		Type: mime,
		Data: data,
	}.Add(ops)
}

// Pos returns the drag position relative to its initial click position.
func (d *Draggable) Pos() f32.Point {
	return d.pos
}
