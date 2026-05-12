//go:build arm64 && !gccgo && !appengine && !arm64_casp
// +build arm64,!gccgo,!appengine,!arm64_casp

package atomic128

func compareAndSwapUint128(*[2]uint64, [2]uint64, [2]uint64) bool
func loadUint128(*[2]uint64) [2]uint64
func storeUint128(*[2]uint64, [2]uint64)
func swapUint128(*[2]uint64, [2]uint64) [2]uint64
func addUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64
func andUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64
func orUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64
func xorUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64

func init() {
	useNative = true
}
