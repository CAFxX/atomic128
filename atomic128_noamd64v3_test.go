//go:build !amd64 || !amd64.v3

package atomic128

import "testing"

func runTests(t *testing.T, fn func(*testing.T)) {
	if hasNative() {
		t.Run("native", fn)
	}
	t.Run("fallback", func(t *testing.T) {
		fallback(t)
		fn(t)
	})
}

func runBenchmarks(b *testing.B, fn func(*testing.PB)) {
	if hasNative() {
		b.Run("native", func(b *testing.B) {
			b.RunParallel(fn)
		})
	}
	b.Run("fallback", func(b *testing.B) {
		fallback(b)
		b.RunParallel(fn)
	})
}

func hasNative() bool {
	return useNativeAmd64
}

func fallback(tb testing.TB) {
	amd64 := useNativeAmd64
	useNativeAmd64 = false
	tb.Cleanup(func() {
		useNativeAmd64 = amd64
	})
}
