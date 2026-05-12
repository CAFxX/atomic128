//go:build amd64.v2 && !amd64.v3 && !gccgo && !appengine

package atomic128

var (
	loadUint128Impl  func(ptr *Uint128) [2]uint64
	storeUint128Impl func(ptr *Uint128, new [2]uint64)
)

func initDispatch(useNativeAmd64, useAVXAmd64 bool) {
	if useAVXAmd64 {
		loadUint128Impl = loadUint128amd64avx
		storeUint128Impl = storeUint128amd64avx
	} else {
		loadUint128Impl = loadUint128amd64
		storeUint128Impl = storeUint128amd64
	}
}

func CompareAndSwapUint128(ptr *Uint128, old, new [2]uint64) bool {
	return compareAndSwapUint128amd64(ptr, old, new)
}

func LoadUint128(ptr *Uint128) [2]uint64 {
	return loadUint128Impl(ptr)
}

func StoreUint128(ptr *Uint128, new [2]uint64) {
	storeUint128Impl(ptr, new)
}

func SwapUint128(ptr *Uint128, new [2]uint64) [2]uint64 {
	return swapUint128amd64(ptr, new)
}

func AddUint128(ptr *Uint128, incr [2]uint64) [2]uint64 {
	return addUint128amd64(ptr, incr)
}

func AndUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return andUint128amd64(ptr, op)
}

func OrUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return orUint128amd64(ptr, op)
}

func XorUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return xorUint128amd64(ptr, op)
}
