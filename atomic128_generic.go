//go:build !amd64 || gccgo || appengine

package atomic128

func initDispatch(bool, bool, bool) {}

// CompareAndSwap performs a 128-bit atomic CAS on ptr.
func (ptr *Uint128) CompareAndSwap(old, new [2]uint64) bool {
	return compareAndSwapUint128Fallback(ptr, old, new)
}

// Load atomically loads the 128 bit value pointed to by ptr.
func (ptr *Uint128) Load() [2]uint64 {
	return loadUint128Fallback(ptr)
}

// Store atomically stores the new value in the 128 bit value pointed to by ptr.
func (ptr *Uint128) Store(new [2]uint64) {
	storeUint128Fallback(ptr, new)
}

// Swap atomically stores the new value with the 128 bit value pointed to by ptr.
func (ptr *Uint128) Swap(new [2]uint64) [2]uint64 {
	return swapUint128Fallback(ptr, new)
}

// Add atomically adds the incr value to the 128 bit value pointed to by ptr.
func (ptr *Uint128) Add(incr [2]uint64) [2]uint64 {
	return addUint128Fallback(ptr, incr)
}

// And atomically performs a bitwise AND of the op value to the 128 bit value pointed to by ptr.
func (ptr *Uint128) And(op [2]uint64) [2]uint64 {
	return andUint128Fallback(ptr, op)
}

// Or atomically performs a bitwise OR of the op value to the 128 bit value pointed to by ptr.
func (ptr *Uint128) Or(op [2]uint64) [2]uint64 {
	return orUint128Fallback(ptr, op)
}

// Xor atomically performs a bitwise XOR of the op value to the 128 bit value pointed to by ptr.
func (ptr *Uint128) Xor(op [2]uint64) [2]uint64 {
	return xorUint128Fallback(ptr, op)
}
