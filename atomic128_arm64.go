// +build arm64,!gccgo,!appengine

package atomic128

import "github.com/klauspost/cpuid/v2"

func compareAndSwapUint128arm64(*[2]uint64, [2]uint64, [2]uint64) bool
func loadUint128arm64(*[2]uint64) [2]uint64
func storeUint128arm64(*[2]uint64, [2]uint64)
func swapUint128arm64(*[2]uint64, [2]uint64) [2]uint64
func addUint128arm64(ptr *[2]uint64, incr [2]uint64) [2]uint64

func init() {
	if !cpuid.CPU.Supports(cpuid.CX16) {
		return
	}
	compareAndSwapUint128 = compareAndSwapUint128arm64
	loadUint128 = loadUint128arm64
	storeUint128 = storeUint128arm64
	swapUint128 = swapUint128arm64
	addUint128 = addUint128arm64
}
