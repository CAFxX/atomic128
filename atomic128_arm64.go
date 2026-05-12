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

func compareAndSwapUint128Native(ptr *Uint128, old, new [2]uint64) bool {
    return compareAndSwapUint128(addr(ptr), old, new)
}

func loadUint128Native(ptr *Uint128) [2]uint64 {
    return loadUint128(addr(ptr))
}

func storeUint128Native(ptr *Uint128, new [2]uint64) {
    storeUint128(addr(ptr), new)
}

func swapUint128Native(ptr *Uint128, new [2]uint64) [2]uint64 {
    return swapUint128(addr(ptr), new)
}

func addUint128Native(ptr *Uint128, incr [2]uint64) [2]uint64 {
    return addUint128(addr(ptr), incr)
}

func andUint128Native(ptr *Uint128, op [2]uint64) [2]uint64 {
    return andUint128(addr(ptr), op)
}

func orUint128Native(ptr *Uint128, op [2]uint64) [2]uint64 {
    return orUint128(addr(ptr), op)
}

func xorUint128Native(ptr *Uint128, op [2]uint64) [2]uint64 {
    return xorUint128(addr(ptr), op)
}

func init() {
	useNative = true
    initDispatch()
}
