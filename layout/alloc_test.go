// SPDX-License-Identifier: Unlicense OR MIT

//go:build !race
// +build !race

package layout

import (
	"image"
	"testing"

	"github.com/Seikaijyu/gio/op"
)

func TestStackAllocs(t *testing.T) {
	var ops op.Ops
	allocs := testing.AllocsPerRun(1, func() {
		ops.Reset()
		gtx := Context{
			Ops: &ops,
		}
		Stack{}.Layout(gtx,
			Stacked(func(gtx Context) Dimensions {
				return Dimensions{Size: image.Point{X: 50, Y: 50}}
			}),
		)
	})
	if allocs != 0 {
		t.Errorf("expected no allocs, got %f", allocs)
	}
}

func TestFlexAllocs(t *testing.T) {
	var ops op.Ops
	allocs := testing.AllocsPerRun(1, func() {
		ops.Reset()
		gtx := Context{
			Ops: &ops,
		}
		Flex{}.Layout(gtx,
			Rigid(func(gtx Context) Dimensions {
				return Dimensions{Size: image.Point{X: 50, Y: 50}}
			}),
		)
	})
	if allocs != 0 {
		t.Errorf("expected no allocs, got %f", allocs)
	}
}
