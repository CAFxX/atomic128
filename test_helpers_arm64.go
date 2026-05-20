//go:build arm64

package atomic128

import (
	"testing"
	"github.com/klauspost/cpuid/v2"
)

func hasNative() bool {
	return true // arm64 always has ldaxp/stlxp natively
}

func fallback(tb testing.TB) {
	lse := cpuid.CPU.Supports(cpuid.ATOMICS)
	initDispatchFallback(true)
	tb.Cleanup(func() {
		initDispatchFallback(false)
		initDispatch(lse)
	})
}
