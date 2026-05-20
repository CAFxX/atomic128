//go:build arm64 && !gccgo && !appengine

package atomic128

var (
	compareAndSwapUint128Impl func(ptr *Uint128, old, new [2]uint64) bool
	loadUint128Impl           func(ptr *Uint128) [2]uint64
	storeUint128Impl          func(ptr *Uint128, new [2]uint64)
	swapUint128Impl           func(ptr *Uint128, new [2]uint64) [2]uint64
	addUint128Impl            func(ptr *Uint128, incr [2]uint64) [2]uint64
	andUint128Impl            func(ptr *Uint128, op [2]uint64) [2]uint64
	orUint128Impl             func(ptr *Uint128, op [2]uint64) [2]uint64
	xorUint128Impl            func(ptr *Uint128, op [2]uint64) [2]uint64
)

func initDispatchFallback(fallback bool) {
	if fallback {
		compareAndSwapUint128Impl = compareAndSwapUint128Fallback
		loadUint128Impl = loadUint128Fallback
		storeUint128Impl = storeUint128Fallback
		swapUint128Impl = swapUint128Fallback
		addUint128Impl = addUint128Fallback
		andUint128Impl = andUint128Fallback
		orUint128Impl = orUint128Fallback
		xorUint128Impl = xorUint128Fallback
	}
}

func initDispatch(useLSE bool) {
	if useLSE {
		compareAndSwapUint128Impl = compareAndSwapUint128arm64caspd
	} else {
		compareAndSwapUint128Impl = compareAndSwapUint128arm64
	}
	loadUint128Impl = loadUint128arm64
	storeUint128Impl = storeUint128arm64
	swapUint128Impl = swapUint128arm64
	addUint128Impl = addUint128arm64
	andUint128Impl = andUint128arm64
	orUint128Impl = orUint128arm64
	xorUint128Impl = xorUint128arm64
}

func CompareAndSwapUint128(ptr *Uint128, old, new [2]uint64) bool {
	return compareAndSwapUint128Impl(ptr, old, new)
}

func LoadUint128(ptr *Uint128) [2]uint64 {
	return loadUint128Impl(ptr)
}

func StoreUint128(ptr *Uint128, new [2]uint64) {
	storeUint128Impl(ptr, new)
}

func SwapUint128(ptr *Uint128, new [2]uint64) [2]uint64 {
	return swapUint128Impl(ptr, new)
}

func AddUint128(ptr *Uint128, incr [2]uint64) [2]uint64 {
	return addUint128Impl(ptr, incr)
}

func AndUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return andUint128Impl(ptr, op)
}

func OrUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return orUint128Impl(ptr, op)
}

func XorUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return xorUint128Impl(ptr, op)
}
