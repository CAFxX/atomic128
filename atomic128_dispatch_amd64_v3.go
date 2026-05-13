//go:build amd64 && amd64.v3 && !gccgo && !appengine
// +build amd64,amd64.v3,!gccgo,!appengine

package atomic128

var (
	CompareAndSwapUint128 func(ptr *Uint128, old, new [2]uint64) bool = compareAndSwapUint128
	LoadUint128           func(ptr *Uint128) [2]uint64                = loadUint128avx
	StoreUint128          func(ptr *Uint128, new [2]uint64)           = storeUint128avx
	SwapUint128           func(ptr *Uint128, new [2]uint64) [2]uint64 = swapUint128
	AddUint128            func(ptr *Uint128, incr [2]uint64) [2]uint64 = addUint128
	AndUint128            func(ptr *Uint128, op [2]uint64) [2]uint64   = andUint128
	OrUint128             func(ptr *Uint128, op [2]uint64) [2]uint64    = orUint128
	XorUint128            func(ptr *Uint128, op [2]uint64) [2]uint64    = xorUint128
)

func initDispatch() {}
