//go:build amd64 && !gccgo && !appengine

package atomic128

import "github.com/klauspost/cpuid/v2"

var (
	useAVXAmd64 bool
)

func compareAndSwapUint128(*Uint128, [2]uint64, [2]uint64) bool
func loadUint128(*Uint128) [2]uint64
func storeUint128(*Uint128, [2]uint64)
func swapUint128(*Uint128, [2]uint64) [2]uint64
func addUint128(ptr *Uint128, incr [2]uint64) [2]uint64
func andUint128(ptr *Uint128, incr [2]uint64) [2]uint64
func orUint128(ptr *Uint128, incr [2]uint64) [2]uint64
func xorUint128(ptr *Uint128, incr [2]uint64) [2]uint64

// AVX variants
func loadUint128avx(*Uint128) [2]uint64
func storeUint128avx(*Uint128, [2]uint64)

func init() {
	useNative = cpuid.CPU.Supports(cpuid.CX16)
	useAVXAmd64 = cpuid.CPU.Supports(cpuid.AVX)
	initDispatch()
}
