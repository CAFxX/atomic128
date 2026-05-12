//go:build amd64 && !amd64.v2 && !amd64.v3 && !gccgo && !appengine

package atomic128

var (
	CompareAndSwapUint128 func(ptr *Uint128, old, new [2]uint64) bool
	LoadUint128           func(ptr *Uint128) [2]uint64
	StoreUint128          func(ptr *Uint128, new [2]uint64)
	SwapUint128           func(ptr *Uint128, new [2]uint64) [2]uint64
	AddUint128            func(ptr *Uint128, incr [2]uint64) [2]uint64
	AndUint128            func(ptr *Uint128, op [2]uint64) [2]uint64
	OrUint128             func(ptr *Uint128, op [2]uint64) [2]uint64
	XorUint128            func(ptr *Uint128, op [2]uint64) [2]uint64
)

func initDispatch() {
	if useNativeAmd64 {
		CompareAndSwapUint128 = compareAndSwapUint128amd64
		if useAVXAmd64 {
			LoadUint128 = loadUint128amd64avx
			StoreUint128 = storeUint128amd64avx
		} else {
			LoadUint128 = loadUint128amd64
			StoreUint128 = storeUint128amd64
		}
		SwapUint128 = swapUint128amd64
		AddUint128 = addUint128amd64
		AndUint128 = andUint128amd64
		OrUint128 = orUint128amd64
		XorUint128 = xorUint128amd64
	} else {
		CompareAndSwapUint128 = compareAndSwapUint128Fallback
		LoadUint128 = loadUint128Fallback
		StoreUint128 = storeUint128Fallback
		SwapUint128 = swapUint128Fallback
		AddUint128 = addUint128Fallback
		AndUint128 = andUint128Fallback
		OrUint128 = orUint128Fallback
		XorUint128 = xorUint128Fallback
	}
}
