//go:build amd64 && !amd64.v3 && !gccgo && !appengine

package atomic128

import "github.com/klauspost/cpuid/v2"

func init() {
	useNativeAmd64 = cpuid.CPU.Supports(cpuid.CX16)
}
