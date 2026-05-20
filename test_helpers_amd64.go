//go:build amd64

package atomic128

import (
	"testing"
	"github.com/klauspost/cpuid/v2"
)

func hasNative() bool {
	return cpuid.CPU.Supports(cpuid.CX16)
}

func fallback(tb testing.TB) {
	native, avx := cpuid.CPU.Supports(cpuid.CX16), cpuid.CPU.Supports(cpuid.AVX)
	initDispatch(false, false)
	tb.Cleanup(func() {
		initDispatch(native, avx)
	})
}
