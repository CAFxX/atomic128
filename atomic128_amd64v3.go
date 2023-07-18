//go:build amd64 && amd64.v3 && !gccgo && !appengine

package atomic128

import "unsafe"

// Uint128 is an opaque container for an atomic uint128.
// Uint128 must not be copied.
// The zero value is a valid value representing [2]uint64{0, 0}.
type Uint128 struct {
	// d is placed first because
	// addr() relies on the 64-bit alignment guarantee: see
	// https://go101.org/article/memory-layout.html and,
	// specifically, https://pkg.go.dev/sync/atomic#pkg-note-BUG.
	// Technically speaking we would need only 3 elements, but in
	// practice by having 4 we should always be aligned (because
	// the struct becomes power-of-2-sized, and power-of-2-sized
	// objects past 16 bytes are always aligned to the size in the
	// standard go runtime) so the if in addr() becomes extremely
	// predictable.
	d [4]uint64
}

// CompareAndSwapUint128 performs a 128-bit atomic CAS on ptr.
// If the memory pointed to by ptr contains the value old, it is set to
// the value new, and true is returned. Otherwise the memory pointed to
// by ptr is unchanged, and false is returned.
// In the old and new values the first of the two elements is the low-order bits.
func CompareAndSwapUint128(ptr *Uint128, old, new [2]uint64) bool {
	return compareAndSwapUint128amd64(addr(ptr), old, new)
}

// LoadUint128 atomically loads the 128 bit value pointed to by ptr.
// In the returned value the first of the two elements is the low-order bits.
func LoadUint128(ptr *Uint128) [2]uint64 {
	return loadUint128amd64(addr(ptr))
}

// StoreUint128 atomically stores the new value in the 128 bit value pointed to by ptr.
// In the new value the first of the two elements is the low-order bits.
func StoreUint128(ptr *Uint128, new [2]uint64) {
	storeUint128amd64(addr(ptr), new)
}

// SwapUint128 atomically stores the new value with the 128 bit value pointed to by ptr,
// and it returns the 128 bit value that was previously pointed to by ptr.
// In the new and returned values the first of the two elements is the low-order bits.
func SwapUint128(ptr *Uint128, new [2]uint64) [2]uint64 {
	return swapUint128amd64(addr(ptr), new)
}

// AddUint128 atomically adds the incr value to the 128 bit value pointed to by ptr,
// and it returns the resulting 128 bit value.
// In the incr and returned values the first of the two elements is the low-order bits.
func AddUint128(ptr *Uint128, incr [2]uint64) [2]uint64 {
	return addUint128amd64(addr(ptr), incr)
}

// AndUint128 atomically performs a bitwise AND of the op value to the 128 bit value pointed to by ptr,
// and it returns the resulting 128 bit value.
// In the op and returned values the first of the two elements is the low-order bits.
func AndUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return andUint128amd64(addr(ptr), op)
}

// OrUint128 atomically performs a bitwise OR of the op value to the 128 bit value pointed to by ptr,
// and it returns the resulting 128 bit value.
// In the op and returned values the first of the two elements is the low-order bits.
func OrUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return orUint128amd64(addr(ptr), op)
}

// XorUint128 atomically performs a bitwise XOR of the op value to the 128 bit value pointed to by ptr,
// and it returns the resulting 128 bit value.
// In the op and returned values the first of the two elements is the low-order bits.
func XorUint128(ptr *Uint128, op [2]uint64) [2]uint64 {
	return xorUint128amd64(addr(ptr), op)
}

func addr(ptr *Uint128) *[2]uint64 {
	if (uintptr)((unsafe.Pointer)(&ptr.d[0]))%16 == 0 {
		return (*[2]uint64)((unsafe.Pointer)(&ptr.d[0]))
	}
	return (*[2]uint64)((unsafe.Pointer)(&ptr.d[1]))
}
