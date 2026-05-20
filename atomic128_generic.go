//go:build (!amd64 && !arm64) || gccgo || appengine

package atomic128

func initDispatch(bool, bool) {}

// CompareAndSwapUint128 performs a 128-bit atomic CAS on ptr.
func CompareAndSwapUint128(ptr *Uint128, old, new [2]uint64) bool {
	return compareAndSwapUint128Fallback(ptr, old, new)
}

// LoadUint128 atomically loads the 128 bit value pointed to by ptr.
func LoadUint128(ptr *Uint128) [2]uint64 {
	return loadUint128Fallback(ptr)
}

// StoreUint128 atomically stores the new value in the 128 bit value pointed to by ptr.
func StoreUint128(ptr *Uint128, new [2]uint64) {
	storeUint128Fallback(ptr, new)
}

// SwapUint128 atomically stores the new value with the 128 bit value pointed to by ptr.
func SwapUint128(ptr *Uint128, new [2]uint64) [2]uint64 {
	return swapUint128Fallback(ptr, new)
}

// AddUint128 atomically adds the incr value to the 128 bit value pointed to by ptr.
func AddUint128(ptr *Uint128, incr [2]uint64) [2]uint64 {
	return addUint128Fallback(ptr, incr)
}

// AndUint128 atomically performs a bitwise AND of the op value to the 128 bit value pointed to by ptr.
func AndUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return andUint128Fallback(ptr, op)
}

// OrUint128 atomically performs a bitwise OR of the op value to the 128 bit value pointed to by ptr.
func OrUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return orUint128Fallback(ptr, op)
}

// XorUint128 atomically performs a bitwise XOR of the op value to the 128 bit value pointed to by ptr.
func XorUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return xorUint128Fallback(ptr, op)
}
