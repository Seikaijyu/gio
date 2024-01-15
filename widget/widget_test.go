// SPDX-License-Identifier: Unlicense OR MIT

package widget_test

import (
	"image"
	"testing"

	"github.com/Seikaijyu/gio/f32"
	"github.com/Seikaijyu/gio/io/pointer"
	"github.com/Seikaijyu/gio/io/router"
	"github.com/Seikaijyu/gio/io/semantic"
	"github.com/Seikaijyu/gio/io/system"
	"github.com/Seikaijyu/gio/layout"
	"github.com/Seikaijyu/gio/op"
	"github.com/Seikaijyu/gio/widget"
)

func TestBool(t *testing.T) {
	var (
		ops op.Ops
		r   router.Router
		b   widget.Bool
	)
	gtx := layout.NewContext(&ops, system.FrameEvent{Queue: &r})
	layout := func() {
		b.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			semantic.CheckBox.Add(gtx.Ops)
			semantic.DescriptionOp("description").Add(gtx.Ops)
			return layout.Dimensions{Size: image.Pt(100, 100)}
		})
	}
	layout()
	r.Frame(gtx.Ops)
	r.Queue(
		pointer.Event{
			Source:   pointer.Touch,
			Kind:     pointer.Press,
			Position: f32.Pt(50, 50),
		},
		pointer.Event{
			Source:   pointer.Touch,
			Kind:     pointer.Release,
			Position: f32.Pt(50, 50),
		},
	)
	ops.Reset()
	layout()
	r.Frame(gtx.Ops)
	tree := r.AppendSemantics(nil)
	n := tree[0].Children[0].Desc
	if n.Description != "description" {
		t.Errorf("unexpected semantic description: %s", n.Description)
	}
	if n.Class != semantic.CheckBox {
		t.Errorf("unexpected semantic class: %v", n.Class)
	}
	if !b.Value || !n.Selected {
		t.Error("click did not select")
	}
}
