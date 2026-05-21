//go:build amd64 && !gccgo && !appengine

package atomic128

import "github.com/klauspost/cpuid/v2"

func compareAndSwapUint128amd64(*Uint128, [2]uint64, [2]uint64) bool
func loadUint128amd64(*Uint128) [2]uint64
func storeUint128amd64(*Uint128, [2]uint64)
func swapUint128amd64(*Uint128, [2]uint64) [2]uint64
func addUint128amd64(ptr *Uint128, incr [2]uint64) [2]uint64
func andUint128amd64(ptr *Uint128, incr [2]uint64) [2]uint64
func orUint128amd64(ptr *Uint128, incr [2]uint64) [2]uint64
func xorUint128amd64(ptr *Uint128, incr [2]uint64) [2]uint64

// AVX variants
func loadUint128amd64avx(*Uint128) [2]uint64
func storeUint128amd64avx(*Uint128, [2]uint64)

func init() {
	initDispatch(cpuid.CPU.Supports(cpuid.CX16), cpuid.CPU.Supports(cpuid.AVX), cpuid.CPU.Supports(cpuid.RTM))
}
