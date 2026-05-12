//go:build amd64 && amd64.v2 && !amd64.v3 && !gccgo && !appengine

package atomic128

var (
	CompareAndSwapUint128 func(ptr *Uint128, old, new [2]uint64) bool = compareAndSwapUint128Fallback
	LoadUint128           func(ptr *Uint128) [2]uint64                = loadUint128Fallback
	StoreUint128          func(ptr *Uint128, new [2]uint64)           = storeUint128Fallback
	SwapUint128           func(ptr *Uint128, new [2]uint64) [2]uint64 = swapUint128Fallback
	AddUint128            func(ptr *Uint128, incr [2]uint64) [2]uint64 = addUint128Fallback
	AndUint128            func(ptr *Uint128, op [2]uint64) [2]uint64   = andUint128Fallback
	OrUint128             func(ptr *Uint128, op [2]uint64) [2]uint64    = orUint128Fallback
	XorUint128            func(ptr *Uint128, op [2]uint64) [2]uint64    = xorUint128Fallback
)

func initDispatch() {
	if useNative {
		CompareAndSwapUint128 = compareAndSwapUint128
		if useAVXAmd64 {
			LoadUint128 = loadUint128avx
			StoreUint128 = storeUint128avx
		} else {
			LoadUint128 = loadUint128
			StoreUint128 = storeUint128
		}
		SwapUint128 = swapUint128
		AddUint128 = addUint128
		AndUint128 = andUint128
		OrUint128 = orUint128
		XorUint128 = xorUint128
	}
}
