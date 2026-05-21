//go:build amd64.v3 && !gccgo && !appengine

package atomic128

var (
	storeUint128Impl func(ptr *Uint128, new [2]uint64)
	swapUint128Impl  func(ptr *Uint128, new [2]uint64) [2]uint64
	addUint128Impl   func(ptr *Uint128, incr [2]uint64) [2]uint64
	andUint128Impl   func(ptr *Uint128, op [2]uint64) [2]uint64
	orUint128Impl    func(ptr *Uint128, op [2]uint64) [2]uint64
	xorUint128Impl   func(ptr *Uint128, op [2]uint64) [2]uint64
)

func initDispatch(useNativeAmd64, useAVXAmd64, useRTMAmd64 bool) {
	if useRTMAmd64 {
		storeUint128Impl = storeUint128amd64rtm
		swapUint128Impl = swapUint128amd64rtm
		addUint128Impl = addUint128amd64rtm
		andUint128Impl = andUint128amd64rtm
		orUint128Impl = orUint128amd64rtm
		xorUint128Impl = xorUint128amd64rtm
	} else {
		storeUint128Impl = storeUint128amd64avx
		swapUint128Impl = swapUint128amd64
		addUint128Impl = addUint128amd64
		andUint128Impl = andUint128amd64
		orUint128Impl = orUint128amd64
		xorUint128Impl = xorUint128amd64
	}
}

func CompareAndSwapUint128(ptr *Uint128, old, new [2]uint64) bool {
	return compareAndSwapUint128amd64(ptr, old, new)
}

func LoadUint128(ptr *Uint128) [2]uint64 {
	return loadUint128amd64avx(ptr)
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
