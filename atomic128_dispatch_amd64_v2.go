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

func (ptr *Uint128) CompareAndSwap(old, new [2]uint64) bool {
	return compareAndSwapUint128amd64(ptr, old, new)
}

func (ptr *Uint128) Load() [2]uint64 {
	return loadUint128Impl(ptr)
}

func (ptr *Uint128) Store(new [2]uint64) {
	storeUint128Impl(ptr, new)
}

func (ptr *Uint128) Swap(new [2]uint64) [2]uint64 {
	return swapUint128amd64(ptr, new)
}

func (ptr *Uint128) Add(incr [2]uint64) [2]uint64 {
	return addUint128amd64(ptr, incr)
}

func (ptr *Uint128) And(op [2]uint64) [2]uint64 {
	return andUint128amd64(ptr, op)
}

func (ptr *Uint128) Or(op [2]uint64) [2]uint64 {
	return orUint128amd64(ptr, op)
}

func (ptr *Uint128) Xor(op [2]uint64) [2]uint64 {
	return xorUint128amd64(ptr, op)
}
