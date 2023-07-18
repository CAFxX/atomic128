//go:build amd64 && amd64.v3 && !gccgo && !appengine

package atomic128

import "testing"

func runTests(t *testing.T, fn func(*testing.T)) {
	t.Run("native", fn)
}

func runBenchmarks(b *testing.B, fn func(*testing.PB)) {
	b.Run("native", func(b *testing.B) {
		b.RunParallel(fn)
	})
}
