//go:build !amd64 || gccgo || appengine
// +build !amd64 gccgo appengine

package atomic128

import "sync"

// Uint128 is an opaque container for an atomic uint128.
// Uint128 must not be copied.
// The zero value is a valid value representing [2]uint64{0, 0}.
type Uint128 struct {
	d [2]uint64
	m sync.Mutex
}

func compareAndSwapUint128Fallback(ptr *Uint128, old, new [2]uint64) bool {
	ptr.m.Lock()
	v := load2(&ptr.d)
	if v != old {
		ptr.m.Unlock()
		return false
	}
	store2(&ptr.d, new)
	ptr.m.Unlock()
	return true
}

func loadUint128Fallback(ptr *Uint128) [2]uint64 {
	ptr.m.Lock()
	v := load2(&ptr.d)
	ptr.m.Unlock()
	return v
}

func storeUint128Fallback(ptr *Uint128, new [2]uint64) {
	ptr.m.Lock()
	store2(&ptr.d, new)
	ptr.m.Unlock()
}

func swapUint128Fallback(ptr *Uint128, new [2]uint64) [2]uint64 {
	ptr.m.Lock()
	old := load2(&ptr.d)
	store2(&ptr.d, new)
	ptr.m.Unlock()
	return old
}

func addUint128Fallback(ptr *Uint128, incr [2]uint64) [2]uint64 {
	ptr.m.Lock()
	v := load2(&ptr.d)
	v[0] += incr[0]
	if v[0] < incr[0] {
		v[1]++
	}
	v[1] += incr[1]
	store2(&ptr.d, v)
	ptr.m.Unlock()
	return v
}

func andUint128Fallback(ptr *Uint128, op [2]uint64) [2]uint64 {
	ptr.m.Lock()
	v := load2(&ptr.d)
	v[0] &= op[0]
	v[1] &= op[1]
	store2(&ptr.d, v)
	ptr.m.Unlock()
	return v
}

func orUint128Fallback(ptr *Uint128, op [2]uint64) [2]uint64 {
	ptr.m.Lock()
	v := load2(&ptr.d)
	v[0] |= op[0]
	v[1] |= op[1]
	store2(&ptr.d, v)
	ptr.m.Unlock()
	return v
}

func xorUint128Fallback(ptr *Uint128, op [2]uint64) [2]uint64 {
	ptr.m.Lock()
	v := load2(&ptr.d)
	v[0] ^= op[0]
	v[1] ^= op[1]
	store2(&ptr.d, v)
	ptr.m.Unlock()
	return v
}
