// SPDX-License-Identifier: Unlicense OR MIT

package gpu

import (
	"testing"

	"github.com/Seikaijyu/gio/internal/f32"
)

func BenchmarkEncodeQuadTo(b *testing.B) {
	var data [vertStride * 4]byte
	for i := 0; i < b.N; i++ {
		v := float32(i)
		encodeQuadTo(data[:], 123,
			f32.Point{X: v, Y: v},
			f32.Point{X: v, Y: v},
			f32.Point{X: v, Y: v},
		)
	}
}
