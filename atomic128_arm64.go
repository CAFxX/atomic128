//go:build arm64 && !gccgo && !appengine

package atomic128

import "github.com/klauspost/cpuid/v2"

func compareAndSwapUint128arm64(*Uint128, [2]uint64, [2]uint64) bool
func loadUint128arm64(*Uint128) [2]uint64
func storeUint128arm64(*Uint128, [2]uint64)
func swapUint128arm64(*Uint128, [2]uint64) [2]uint64
func addUint128arm64(ptr *Uint128, incr [2]uint64) [2]uint64
func andUint128arm64(ptr *Uint128, incr [2]uint64) [2]uint64
func orUint128arm64(ptr *Uint128, incr [2]uint64) [2]uint64
func xorUint128arm64(ptr *Uint128, incr [2]uint64) [2]uint64

// LSE variant for CAS (CASPD)
func compareAndSwapUint128arm64caspd(*Uint128, [2]uint64, [2]uint64) bool

func init() {
	// ATOMICS maps to LSE (Large System Extensions) on arm64 which includes CASPD.
	initDispatch(cpuid.CPU.Supports(cpuid.ATOMICS))
}
