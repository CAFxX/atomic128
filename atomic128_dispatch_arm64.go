//go:build arm64 && !gccgo && !appengine

package atomic128

var (
	CompareAndSwapUint128 func(ptr *Uint128, old, new [2]uint64) bool = compareAndSwapUint128_arm64
	LoadUint128           func(ptr *Uint128) [2]uint64                = loadUint128_arm64
	StoreUint128          func(ptr *Uint128, new [2]uint64)           = storeUint128_arm64
	SwapUint128           func(ptr *Uint128, new [2]uint64) [2]uint64 = swapUint128_arm64
	AddUint128            func(ptr *Uint128, incr [2]uint64) [2]uint64 = addUint128_arm64
	AndUint128            func(ptr *Uint128, op [2]uint64) [2]uint64   = andUint128_arm64
	OrUint128             func(ptr *Uint128, op [2]uint64) [2]uint64    = orUint128_arm64
	XorUint128            func(ptr *Uint128, op [2]uint64) [2]uint64    = xorUint128_arm64
)

func compareAndSwapUint128_arm64(ptr *Uint128, old, new [2]uint64) bool {
    if useNative {
        return compareAndSwapUint128(addr(ptr), old, new)
    }
    return compareAndSwapUint128Fallback(ptr, old, new)
}
func loadUint128_arm64(ptr *Uint128) [2]uint64 {
    if useNative {
        return loadUint128(addr(ptr))
    }
    return loadUint128Fallback(ptr)
}
func storeUint128_arm64(ptr *Uint128, new [2]uint64) {
    if useNative {
        storeUint128(addr(ptr), new)
        return
    }
    storeUint128Fallback(ptr, new)
}
func swapUint128_arm64(ptr *Uint128, new [2]uint64) [2]uint64 {
    if useNative {
        return swapUint128(addr(ptr), new)
    }
    return swapUint128Fallback(ptr, new)
}
func addUint128_arm64(ptr *Uint128, incr [2]uint64) [2]uint64 {
    if useNative {
        return addUint128(addr(ptr), incr)
    }
    return addUint128Fallback(ptr, incr)
}
func andUint128_arm64(ptr *Uint128, op [2]uint64) [2]uint64 {
    if useNative {
        return andUint128(addr(ptr), op)
    }
    return andUint128Fallback(ptr, op)
}
func orUint128_arm64(ptr *Uint128, op [2]uint64) [2]uint64 {
    if useNative {
        return orUint128(addr(ptr), op)
    }
    return orUint128Fallback(ptr, op)
}
func xorUint128_arm64(ptr *Uint128, op [2]uint64) [2]uint64 {
    if useNative {
        return xorUint128(addr(ptr), op)
    }
    return xorUint128Fallback(ptr, op)
}
