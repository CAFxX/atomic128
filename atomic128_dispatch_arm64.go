//go:build arm64 && !gccgo && !appengine
// +build arm64,!gccgo,!appengine

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
	if useNative {
		CompareAndSwapUint128 = compareAndSwapUint128Native
		LoadUint128 = loadUint128Native
		StoreUint128 = storeUint128Native
		SwapUint128 = swapUint128Native
		AddUint128 = addUint128Native
		AndUint128 = andUint128Native
		OrUint128 = orUint128Native
		XorUint128 = xorUint128Native
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
